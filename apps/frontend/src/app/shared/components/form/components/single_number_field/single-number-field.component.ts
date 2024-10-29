import { Component, inject, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  ControlContainer,
  FormControl,
  FormGroup,
  ValidatorFn,
} from '@angular/forms';

@Component({
  selector: 'app-single-number-field',
  templateUrl: './single-number-field.component.html',
  styleUrl: './single-number-field.component.scss',
  viewProviders: [
    {
      provide: ControlContainer,
      useFactory: () => inject(ControlContainer, { skipSelf: true }),
    },
  ],
})
export class SingleNumberFieldComponent {
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
}
