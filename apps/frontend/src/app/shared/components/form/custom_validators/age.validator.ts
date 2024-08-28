// src/app/validators/age-validator.ts
import { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

export const ageValidator = (minAge: number): ValidatorFn => {
  return (control: AbstractControl): ValidationErrors | null => {
    const birthDate = new Date(control.value);
    const today = new Date();
    const age = today.getFullYear() - birthDate.getFullYear();
    const monthDiff = today.getMonth() - birthDate.getMonth();

    if (
      monthDiff < 0 ||
      (monthDiff === 0 && today.getDate() < birthDate.getDate())
    ) {
      return { tooYoung: true };
    }

    return age >= minAge ? null : { tooYoung: true };
  };
};
