import {
  AfterViewInit,
  ChangeDetectorRef,
  Component,
  EventEmitter,
  Input,
  Output,
} from '@angular/core';
import { FormArray, FormControl, FormGroup } from '@angular/forms';
import {
  AllInputFieldsTypeWithLabel,
  InputFieldTextTypes,
} from 'shared/components/form';

@Component({
  selector: 'app-basic-form',
  templateUrl: './basic-form.component.html',
  styleUrl: './basic-form.component.scss',
})
export class BasicFormComponent implements AfterViewInit {
  form = new FormGroup({});
  @Input({ required: true }) formFields: AllInputFieldsTypeWithLabel[] = [];
  @Input() isLoading: boolean = false;
  @Output() onSubmit = new EventEmitter();

  public get InputFieldTextTypes(): typeof InputFieldTextTypes {
    return InputFieldTextTypes;
  }
  constructor(private readonly cdr: ChangeDetectorRef) {}
  ngAfterViewInit(): void {
    this.cdr.detectChanges();
  }

  public clearForm() {
    this._clearForm(this.form);
    this.cdr.detectChanges();
  }

  public isFormValuesEmpty(): boolean {
    return this._isFormValuesEmpty(this.form);
  }

  private _isFormValuesEmpty(form: FormGroup | FormArray): boolean {
    for (const key of Object.keys(form.controls)) {
      const control = form.get(key);
      if (control instanceof FormControl) {
        if (
          control.value !== null &&
          control.value !== '' &&
          control.value !== undefined
        ) {
          return false;
        }
      } else if (control instanceof FormGroup || control instanceof FormArray) {
        if (this._isFormValuesEmpty(control)) {
          return false;
        }
      }
    }
    return true;
  }

  private _clearForm(form: FormGroup): void {
    Object.keys(form.controls).forEach((key) => {
      const control = form.get(key);

      if (control) {
        // Null check to ensure control exists
        if (control instanceof FormGroup) {
          this._clearForm(control); // Recursive call for nested FormGroups
        } else if (control instanceof FormArray) {
          control.clear(); // Clear the FormArray
        } else {
          control.setValue(null); // Reset individual control value
          control.markAsPristine();
          control.markAsUntouched();
        }
      }
    });
  }

  public submit() {
    this.onSubmit.emit(this.form.value);
  }
}
