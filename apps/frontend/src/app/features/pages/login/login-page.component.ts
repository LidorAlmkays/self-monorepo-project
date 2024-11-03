import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { formFields } from './login.form';
import { AllInputFieldsTypeWithLabel, FormModule } from 'shared/components';
import { LoginStore } from './component_store/login-page.store';
import { UserLoginModel, UserRegisterModel } from 'shared/models';
import { UserService } from 'shared/services/user.services';
import { CardModule } from 'primeng/card';
import { ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-login-page',
  standalone: true,
  providers: [LoginStore, UserService],
  imports: [CommonModule, CardModule, FormModule, ReactiveFormsModule],

  templateUrl: './login-page.component.html',
  styleUrl: './login-page.component.css',
})
export class LoginPageComponent {
  formFields: AllInputFieldsTypeWithLabel[] = formFields;
  loginStoreVm$;

  constructor(private readonly loginStore: LoginStore) {
    this.loginStoreVm$ = this.loginStore.vm$;
  }

  public onSubmit(event: UserLoginModel) {
    this.loginStore.loginUser(event);
  }
}
