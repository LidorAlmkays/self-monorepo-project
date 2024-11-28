import { Route } from '@angular/router';
import { RegisterPageComponent } from './features/pages/register/register-page.component';
import { LoginPageComponent } from './features/pages/login/login-page.component';
import { DownloadYoutubeVideoPageComponent } from './features/pages/download_youtube_video/download-youtube-video-page.component';
import { HomePageComponent } from './features/pages/home/home-page.component';
import { AuthenticateUserTokenGuard } from './core/guards/authenticate-user-token.guard';

export const AppPaths = {
  register: () => 'register',
  login: () => 'login',
  downloadYoutubeVideo: () => 'download-youtube-video',
  home: () => 'home',
};

export const appRoutes: Route[] = [
  {
    path: AppPaths.register(),
    component: RegisterPageComponent,
  },
  {
    path: AppPaths.login(),
    component: LoginPageComponent,
  },
  {
    path: AppPaths.downloadYoutubeVideo(),
    canDeactivate: [AuthenticateUserTokenGuard],
    component: DownloadYoutubeVideoPageComponent,
  },
  {
    path: AppPaths.home(),
    canDeactivate: [AuthenticateUserTokenGuard],
    component: HomePageComponent,
  },
];
