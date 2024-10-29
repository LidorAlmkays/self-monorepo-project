import { Component, inject, Input, OnDestroy, OnInit } from '@angular/core';
import {
  ControlContainer,
  FormControl,
  FormGroup,
  ValidatorFn,
} from '@angular/forms';

@Component({
  selector: 'app-single-input-field',
  templateUrl: './single-input-field.component.html',
  styleUrl: './single-input-field.component.scss',
  viewProviders: [
    {
      provide: ControlContainer,
      useFactory: () => inject(ControlContainer, { skipSelf: true }),
    },
  ],
})
export class SingleInputFieldComponent implements OnInit, OnDestroy {
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
