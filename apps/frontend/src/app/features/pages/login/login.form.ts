import { Validators } from '@angular/forms';
import {
  AllInputFieldsTypeWithLabel,
  InputFieldTextTypes,
} from 'shared/components';

export const formFields: AllInputFieldsTypeWithLabel[] = [
  {
    controlKey: 'username',
    label: 'UserName:',
    inputType: InputFieldTextTypes.text,
    validators: [Validators.required, Validators.maxLength(64)],
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
