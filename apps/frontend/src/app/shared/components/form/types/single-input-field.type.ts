import { ValidatorFn } from '@angular/forms';
import {
  InputFieldTextTypes,
  SinglePasswordFieldOptionsModel,
} from 'shared/components/form';

type WithLabel<T> = T & { label: string };

export type SingleInputFieldType<T extends InputFieldTextTypes> = {
  controlKey: string;
  inputType: T;
  validators?: ValidatorFn[];
  options?: T extends keyof OptionsTypeMap ? OptionsTypeMap[T] : never;
  placeholder?: string;
};

type OptionsTypeMap = {
  [InputFieldTextTypes.password]: SinglePasswordFieldOptionsModel;
  [InputFieldTextTypes.text]: never;
  [InputFieldTextTypes.number]: never;
  [InputFieldTextTypes.date]: never;
};

export type AllInputFieldsTypeWithLabel =
  | WithLabel<SingleInputFieldType<InputFieldTextTypes.password>>
  | WithLabel<SingleInputFieldType<InputFieldTextTypes.number>>
  | WithLabel<SingleInputFieldType<InputFieldTextTypes.date>>
  | WithLabel<SingleInputFieldType<InputFieldTextTypes.text>>;
