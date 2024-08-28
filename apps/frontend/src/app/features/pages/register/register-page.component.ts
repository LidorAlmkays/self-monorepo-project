import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { CardModule } from 'primeng/card';
import {
  AllInputFieldsTypeWithLabel,
  FormModule,
} from 'shared/components/form';
import { formFields } from './register-format.config';
import { RegisterStore } from './component_store/register-page.store';
import { UserService } from 'shared/components/services/user.service';
import { UserModel } from 'shared/models';

@Component({
  selector: 'page-register-page',
  standalone: true,
  providers: [RegisterStore, UserService],
  imports: [CommonModule, ReactiveFormsModule, FormModule, CardModule],
  templateUrl: './register-page.component.html',
  styleUrl: './register-page.component.scss',
})
export class RegisterPageComponent {
  formFields: AllInputFieldsTypeWithLabel[] = formFields;
  registerStoreVm$;
  constructor(private readonly registerStore: RegisterStore) {
    this.registerStoreVm$ = this.registerStore.vm$;
  }

  public onSubmit(event: UserModel) {
    this.registerStore.registerUser(event);
  }
}
