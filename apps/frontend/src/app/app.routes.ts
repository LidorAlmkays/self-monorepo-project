import { Route } from '@angular/router';
import { RegisterPageComponent } from './features/pages/register/register-page.component';
import { LoginPageComponent } from './features/pages/login/login-page.component';
import { DownloadYoutubeVideoPageComponent } from './features/pages/download_youtube_video/download-youtube-video-page.component';
import { HomePageComponent } from './features/pages/home/home-page.component';
import { FormDiartyGuard } from './core/guards/form.diarty.guard';

export const AppPaths = {
  register: () => 'register',
  login: () => 'login',
  downloadYoutubeVideo: () => 'download-youtube-video',
  home: () => 'home',
};

export const appRoutes: Route[] = [
  {
    path: AppPaths.register(),
    canDeactivate: [FormDiartyGuard],
    component: RegisterPageComponent,
  },
  {
    path: AppPaths.login(),
    canDeactivate: [FormDiartyGuard],
    component: LoginPageComponent,
  },
  {
    path: AppPaths.downloadYoutubeVideo(),
    component: DownloadYoutubeVideoPageComponent,
  },
  { path: AppPaths.home(), component: HomePageComponent },
];
