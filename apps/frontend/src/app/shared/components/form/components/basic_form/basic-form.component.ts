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

  isFormValuesEmpty(): boolean {
    return this.doesFormHasValueInsideIt(this.form);
  }

  doesFormHasValueInsideIt(form: FormGroup | FormArray): boolean {
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
        if (this.doesFormHasValueInsideIt(control)) {
          return false;
        }
      }
    }
    return true;
  }

  submit() {
    this.onSubmit.emit(this.form.value);
  }
}
