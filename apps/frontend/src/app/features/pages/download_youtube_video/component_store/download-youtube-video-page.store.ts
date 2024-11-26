import { Injectable } from '@angular/core';
import { ComponentStore } from '@ngrx/component-store';
import { tapResponse } from '@ngrx/operators';
import { Observable, switchMap } from 'rxjs';
import { DownloadYoutubeVideoModel } from 'shared/models/youtube-download-video.model';
import { YoutubeVideoService } from 'shared/services/youtube-video.serivces';

export interface DownloadingVideoState {
  isLoading: boolean;
}

@Injectable()
export class DownloadYoutubeVideoStore extends ComponentStore<DownloadingVideoState> {
  private readonly isLoading$: Observable<boolean> = this.select(
    (state) => state.isLoading
  );
  readonly vm$ = this.select({
    isLoading: this.isLoading$,
  });

  private readonly setIsLoading = this.updater((state, isLoading: boolean) => {
    const newState: DownloadingVideoState = {
      ...state,
      isLoading,
    };
    return newState;
  });

  constructor(private readonly youtubeVideoService: YoutubeVideoService) {
    super({ isLoading: false });
  }

  readonly downloadVideo = this.effect(
    (trigger$: Observable<DownloadYoutubeVideoModel>) => {
      return trigger$.pipe(
        switchMap((downloadYoutubeVideoData) => {
          this.setIsLoading(true);
          return this.youtubeVideoService
            .downloadVideo(downloadYoutubeVideoData)
            .pipe(
              tapResponse(
                (response) => {
                  this.setIsLoading(false);
                  // Handle successful response here
                },
                (error) => {
                  this.setIsLoading(false);
                  // Handle error here
                }
              )
            );
        })
      );
    }
  );
}
