import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'apps/frontend/src/environments/environment';
import { catchError, first, map, Observable, tap, throwError } from 'rxjs';
import { DownloadYoutubeVideoModel } from 'shared/models/youtube-download-video.model';

@Injectable()
export class YoutubeVideoService {
  private readonly youtubeVideoGatewayUrl =
    environment.gateway + '/youtube-video';

  constructor() {}

  downloadVideo(model: DownloadYoutubeVideoModel): Observable<void> {
    return new Observable<void>((observer) => {
      fetch(this.youtubeVideoGatewayUrl + '/download', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(model),
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }

          // Extract filename from Content-Disposition header
          const contentDisposition = response.headers.get(
            'Content-Disposition'
          );

          const filename = contentDisposition
            ? this.getFilenameFromContentDisposition(contentDisposition)
            : 'downloaded-video.mp4'; // Default filename if none found
          const reader = response.body?.getReader();
          if (!reader) {
            throw new Error('Failed to get reader from response body');
          }
          return { reader, filename };
        })
        .then(({ reader, filename }) => {
          const stream = new ReadableStream({
            start(controller) {
              function pump() {
                reader.read().then(({ done, value }) => {
                  if (done) {
                    controller.close();
                    return;
                  }
                  controller.enqueue(value);
                  pump();
                });
              }
              pump();
            },
          });

          return new Response(stream)
            .blob()
            .then((blob) => ({ blob, filename }));
        })
        .then(({ blob, filename }) => {
          const downloadUrl = window.URL.createObjectURL(blob);
          const anchor = document.createElement('a');
          anchor.href = downloadUrl;
          anchor.download = filename;
          anchor.click();
          window.URL.revokeObjectURL(downloadUrl);

          observer.next(); // Signal completion
          observer.complete(); // Mark the observable as complete
        })
        .catch((error) => {
          observer.error(
            new Error('Failed to download video: ' + error.message)
          ); // Emit an error if anything goes wrong
        });
    });
  }

  private getFilenameFromContentDisposition(
    contentDisposition: string
  ): string {
    // Match and capture only the filename portion in the Content-Disposition header
    const matches = /filename[^;=\n]*=\s*(?:(['"]).*?\1|[^;\n]*)/.exec(
      contentDisposition
    );
    return matches
      ? matches[0].replace(/(filename=|['"])/g, '').trim()
      : 'downloaded-video.mp4';
  }
}
