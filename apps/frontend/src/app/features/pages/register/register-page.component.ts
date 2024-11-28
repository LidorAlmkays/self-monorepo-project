import { Component, ViewChild, viewChild } from '@angular/core';
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

@Component({
  selector: 'page-register-page',
  standalone: true,
  providers: [RegisterStore, UserService],
  imports: [CommonModule, ReactiveFormsModule, FormModule, CardModule],
  templateUrl: './register-page.component.html',
  styleUrl: './register-page.component.scss',
})
export class RegisterPageComponent implements IFormEmptyDataSafe {
  formFields: AllInputFieldsTypeWithLabel[] = formFields;
  registerStoreVm$;
  @ViewChild(BasicFormComponent) form!: BasicFormComponent;

  constructor(private readonly registerStore: RegisterStore) {
    this.registerStoreVm$ = this.registerStore.vm$;
  }

  public onSubmit(event: UserRegisterModel) {
    this.registerStore.registerUser(event);
  }

  isFormEmpty(): boolean {
    return this.form.isFormValuesEmpty();
  }
}
