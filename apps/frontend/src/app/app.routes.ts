import { Route } from '@angular/router';
import { RegisterPageComponent } from './features/pages/register/register-page.component';
import { LoginPageComponent } from './features/pages/login/login-page.component';
import { DownloadYoutubeVideoPageComponent } from './features/pages/download_youtube_video/download-youtube-video-page.component';

export const appRoutes: Route[] = [
  {
    path: 'register',
    component: RegisterPageComponent,
  },
  {
    path: 'login',
    component: LoginPageComponent,
  },
  {
    path: 'download-youtube-video',
    component: DownloadYoutubeVideoPageComponent,
  },
];
