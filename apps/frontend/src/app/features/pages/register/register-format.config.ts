import { Validators } from '@angular/forms';
import {
  AllInputFieldsTypeWithLabel,
  InputFieldTextTypes,
} from 'shared/components';
import { ageValidator } from 'shared/components/form/custom_validators';

export const formFields: AllInputFieldsTypeWithLabel[] = [
  {
    controlKey: 'Name',
    label: 'Name:',
    inputType: InputFieldTextTypes.text,
    validators: [Validators.required, Validators.maxLength(64)],
  },
  {
    controlKey: 'username',
    label: 'UserName:',
    inputType: InputFieldTextTypes.text,
    validators: [Validators.required, Validators.maxLength(64)],
  },
  {
    controlKey: 'email',
    label: 'Email:',
    inputType: InputFieldTextTypes.text,
    validators: [Validators.required, Validators.email],
  },
  {
    controlKey: 'age',
    label: 'Age:',
    inputType: InputFieldTextTypes.date,
    validators: [Validators.required, ageValidator(18)],
  },
  {
    controlKey: 'password',
    label: 'Password:',
    inputType: InputFieldTextTypes.password,
    validators: [
      Validators.required,
      Validators.minLength(8),
      Validators.maxLength(64),
    ],
    options: {
      feedback: true,
      toggleMask: true,
    },
  },
];
