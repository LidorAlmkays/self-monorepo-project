import { NgModule } from '@angular/core';
import { RegisterPageComponent } from './register/register-page.component';
import { LoginPageComponent } from './login/login-page.component';

@NgModule({
  imports: [RegisterPageComponent, LoginPageComponent],
  exports: [RegisterPageComponent, LoginPageComponent],
})
export class PagesModule {}
