import { NgModule } from '@angular/core';
import { RegisterPageComponent } from './register/register-page.component';
import { LoginPageComponent } from './login/login-page.component';
import { DownloadYoutubeVideoPageComponent } from './download_youtube_video/download-youtube-video-page.component';

@NgModule({
  imports: [
    RegisterPageComponent,
    DownloadYoutubeVideoPageComponent,
    LoginPageComponent,
  ],
  exports: [
    RegisterPageComponent,
    DownloadYoutubeVideoPageComponent,
    LoginPageComponent,
  ],
})
export class PagesModule {}
