import { Route } from '@angular/router';
import { RegisterPageComponent } from './features/pages/register/register-page.component';

export const appRoutes: Route[] = [
  {
    path: 'register',
    component: RegisterPageComponent,
  },
];
