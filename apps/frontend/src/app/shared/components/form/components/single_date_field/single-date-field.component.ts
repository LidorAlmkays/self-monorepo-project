import { Component, inject, Input } from '@angular/core';
import {
  ControlContainer,
  FormControl,
  FormGroup,
  ValidatorFn,
} from '@angular/forms';

@Component({
  selector: 'app-single-date-field',
  templateUrl: './single-date-field.component.html',
  styleUrl: './single-date-field.component.scss',
  viewProviders: [
    {
      provide: ControlContainer,
      useFactory: () => inject(ControlContainer, { skipSelf: true }),
    },
  ],
})
export class SingleDateFieldComponent {
  @Input({ required: true }) controlKey = '';
  @Input() placeholder?: string = '';
  @Input() validators?: ValidatorFn[] = [];

  private parentContainer = inject(ControlContainer);

  get parentFormGroup() {
    return this.parentContainer.control as FormGroup;
  }

  ngOnInit() {
    this.parentFormGroup.addControl(
      this.controlKey,
      new FormControl(null, this.validators)
    );
  }

  ngOnDestroy() {
    this.parentFormGroup.removeControl(this.controlKey);
  }

  onDateSelect(event: Date) {
    this.parentFormGroup.get(this.controlKey)?.updateValueAndValidity();
  }
}
