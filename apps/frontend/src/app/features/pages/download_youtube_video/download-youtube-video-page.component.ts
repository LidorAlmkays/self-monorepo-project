import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { formFields } from './download-youtube-video.form';
import { AllInputFieldsTypeWithLabel, FormModule } from 'shared/components';
import { UserLoginModel, UserRegisterModel } from 'shared/models';
import { UserService } from 'shared/services/user.services';
import { CardModule } from 'primeng/card';
import { ReactiveFormsModule } from '@angular/forms';
import { DownloadYoutubeVideoStore } from './component_store/download-youtube-video-page.store';
import { DownloadYoutubeVideoModel } from 'shared/models/youtube-download-video.model';
import { YoutubeVideoService } from 'shared/services/youtube-video.serivces';

@Component({
  selector: 'download-youtube-video-page',
  standalone: true,
  providers: [DownloadYoutubeVideoStore, YoutubeVideoService],
  imports: [CommonModule, CardModule, FormModule, ReactiveFormsModule],

  templateUrl: './download-youtube-video-page.component.html',
  styleUrl: './download-youtube-video-page.component.css',
})
export class DownloadYoutubeVideoPageComponent {
  public downloadYoutubeVideoStoreVm$;
  constructor(
    private readonly downloadYoutubeVideoStore: DownloadYoutubeVideoStore
  ) {
    this.downloadYoutubeVideoStoreVm$ = this.downloadYoutubeVideoStore.vm$;
  }
  formFields: AllInputFieldsTypeWithLabel[] = formFields;

  public onSubmit(event: DownloadYoutubeVideoModel) {
    this.downloadYoutubeVideoStore.downloadVideo(event);
  }
}
