import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { CardModule } from 'primeng/card';
import {
  AllInputFieldsTypeWithLabel,
  FormModule,
} from 'shared/components/form';
import { formFields } from './register-format.config';

@Component({
  selector: 'page-register-page',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, FormModule, CardModule],
  templateUrl: './register-page.component.html',
  styleUrl: './register-page.component.scss',
})
export class RegisterPageComponent {
  formFields: AllInputFieldsTypeWithLabel[] = formFields;

  public onSubmit(event: any) {
    console.log(event);
  }
}
