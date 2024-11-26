import { Validators } from '@angular/forms';
import {
  AllInputFieldsTypeWithLabel,
  InputFieldTextTypes,
} from 'shared/components';

export const formFields: AllInputFieldsTypeWithLabel[] = [
  {
    controlKey: 'url',
    label: 'Url:',
    inputType: InputFieldTextTypes.text,
    validators: [Validators.required],
    placeholder: 'Enter url',
  },
];
