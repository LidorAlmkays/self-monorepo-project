import { NgModule } from '@angular/core';
import { SingleInputFieldComponent } from './components/single_input_field/single-input-field.component';
import { InputTextModule } from 'primeng/inputtext';
import { ReactiveFormsModule } from '@angular/forms';
import { PasswordModule } from 'primeng/password';
import { CommonModule } from '@angular/common';
import { DividerModule } from 'primeng/divider';
import { SinglePasswordFieldComponent } from './components/single_password_field/single-password-field.component';
import { BasicFormComponent } from './components/basic_form/basic-form.component';
import { ButtonModule } from 'primeng/button';
import { SingleNumberFieldComponent } from './components/single_number_field/single-number-field.component';
import { InputNumberModule } from 'primeng/inputnumber';
import { SingleDateFieldComponent } from './components/single_date_field/single-date-field.component';
import { CalendarModule } from 'primeng/calendar';

@NgModule({
  declarations: [
    SingleInputFieldComponent,
    SinglePasswordFieldComponent,
    BasicFormComponent,
    SingleNumberFieldComponent,
    SingleDateFieldComponent,
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    InputTextModule,
    PasswordModule,
    DividerModule,
    ButtonModule,
    InputNumberModule,
    CalendarModule,
  ],
  exports: [
    SingleInputFieldComponent,
    SinglePasswordFieldComponent,
    BasicFormComponent,
    SingleNumberFieldComponent,
    SingleDateFieldComponent,
  ],
})
export class FormModule {}
