import { Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
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
import { Subscription } from 'rxjs';
import { Store } from '@ngrx/store';
import { Router } from '@angular/router';
import { userFeature } from '../../ngrx_store/user/user.state';
import { AppPaths } from '../../../app.routes';

@Component({
  selector: 'app-login-page',
  standalone: true,
  providers: [LoginStore, UserService],
  imports: [CommonModule, CardModule, FormModule, ReactiveFormsModule],

  templateUrl: './login-page.component.html',
  styleUrl: './login-page.component.css',
})
export class LoginPageComponent
  implements IFormEmptyDataSafe, OnInit, OnDestroy
{
  formFields: AllInputFieldsTypeWithLabel[] = formFields;
  loginStoreVm$;
  @ViewChild(BasicFormComponent) form!: BasicFormComponent;
  private subscription: Subscription = new Subscription();

  constructor(
    private readonly loginStore: LoginStore,
    private readonly store: Store,
    private readonly router: Router
  ) {
    this.loginStoreVm$ = this.loginStore.vm$;
  }
  ngOnInit(): void {
    this.subscription.add(
      this.store.select(userFeature.selectLoggedIn).subscribe((loggedIn) => {
        if (loggedIn) {
          this.form.clearForm();
          this.router.navigate([AppPaths.home()]);
        }
      })
    );
  }
  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }

  isFormEmpty(): boolean {
    return this.form.isFormValuesEmpty();
  }

  public onSubmit(event: UserLoginModel) {
    this.loginStore.loginUser(event);
  }
}
