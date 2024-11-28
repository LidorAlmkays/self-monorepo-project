import { Component, ViewChild } from '@angular/core';
import { CommonModule } from '@angular/common';
import { formFields } from './login.form';
import { AllInputFieldsTypeWithLabel, FormModule } from 'shared/components';
import { LoginStore } from './component_store/login-page.store';
import { UserLoginModel, UserRegisterModel } from 'shared/models';
import { UserService } from 'shared/services/user.services';
import { CardModule } from 'primeng/card';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';
import { IFormEmptyDataSafe } from 'shared/interfaces';
import { BasicFormComponent } from 'shared/components/form/components/basic_form/basic-form.component';

@Component({
  selector: 'app-login-page',
  standalone: true,
  providers: [LoginStore, UserService],
  imports: [CommonModule, CardModule, FormModule, ReactiveFormsModule],

  templateUrl: './login-page.component.html',
  styleUrl: './login-page.component.css',
})
export class LoginPageComponent implements IFormEmptyDataSafe {
  formFields: AllInputFieldsTypeWithLabel[] = formFields;
  loginStoreVm$;
  @ViewChild(BasicFormComponent) form!: BasicFormComponent;

  constructor(private readonly loginStore: LoginStore) {
    this.loginStoreVm$ = this.loginStore.vm$;
  }

  isFormEmpty(): boolean {
    return this.form.isFormValuesEmpty();
  }

  public onSubmit(event: UserLoginModel) {
    this.loginStore.loginUser(event);
  }
}
