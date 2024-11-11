import { Validators } from '@angular/forms';
import {
  AllInputFieldsTypeWithLabel,
  InputFieldTextTypes,
} from 'shared/components';

export const formFields: AllInputFieldsTypeWithLabel[] = [
  {
    controlKey: 'email',
    label: 'Email:',
    inputType: InputFieldTextTypes.text,
    validators: [Validators.required, Validators.email],
  },
  {
    controlKey: 'password',
    label: 'Password:',
    inputType: InputFieldTextTypes.password,
    validators: [Validators.required],
    options: {
      toggleMask: true,
    },
  },
];
