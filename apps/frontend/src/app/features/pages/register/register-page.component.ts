import {
  Component,
  OnDestroy,
  OnInit,
  ViewChild,
  viewChild,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';
import { CardModule } from 'primeng/card';
import {
  AllInputFieldsTypeWithLabel,
  FormModule,
} from 'shared/components/form';
import { formFields } from './register-format.config';
import { RegisterStore } from './component_store/register-page.store';
import { UserRegisterModel } from 'shared/models';
import { UserService } from 'shared/services/user.services';
import { IFormEmptyDataSafe } from 'shared/interfaces';
import { BasicFormComponent } from 'shared/components/form/components/basic_form/basic-form.component';
import { Store } from '@ngrx/store';
import { userFeature } from '../../ngrx_store/user/user.state';
import { AppPaths } from '../../../app.routes';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';

@Component({
  selector: 'page-register-page',
  standalone: true,
  providers: [RegisterStore, UserService],
  imports: [CommonModule, ReactiveFormsModule, FormModule, CardModule],
  templateUrl: './register-page.component.html',
  styleUrl: './register-page.component.scss',
})
export class RegisterPageComponent
  implements IFormEmptyDataSafe, OnInit, OnDestroy
{
  formFields: AllInputFieldsTypeWithLabel[] = formFields;
  registerStoreVm$;
  @ViewChild(BasicFormComponent) form!: BasicFormComponent;
  private subscription: Subscription = new Subscription();

  constructor(
    private readonly registerStore: RegisterStore,
    private readonly store: Store,
    private readonly router: Router
  ) {
    this.registerStoreVm$ = this.registerStore.vm$;
  }
  ngOnInit(): void {
    this.subscription.add(
      this.store.select(userFeature.selectLoggedIn).subscribe((loggedIn) => {
        if (loggedIn) {
          this.router.navigate([AppPaths.home()]);
        }
      })
    );
  }
  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }

  public onSubmit(event: UserRegisterModel) {
    this.registerStore.registerUser(event);
  }

  isFormEmpty(): boolean {
    return this.form.isFormValuesEmpty();
  }
}
