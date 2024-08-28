import { Component, inject, Input, OnDestroy, OnInit } from '@angular/core';
import {
  ControlContainer,
  FormControl,
  FormGroup,
  ValidatorFn,
} from '@angular/forms';
import { SinglePasswordFieldOptionsModel } from 'shared/components/form';
@Component({
  selector: 'app-single-password-field',
  templateUrl: './single-password-field.component.html',
  styleUrl: './single-password-field.component.scss',
  viewProviders: [
    {
      provide: ControlContainer,
      useFactory: () => inject(ControlContainer, { skipSelf: true }),
    },
  ],
})
export class SinglePasswordFieldComponent implements OnInit, OnDestroy {
  @Input({ required: true }) controlKey = '';
  @Input() placeholder?: string = '';
  @Input() validators?: ValidatorFn[] = [];

  @Input() options?: SinglePasswordFieldOptionsModel = {
    feedback: false,
    toggleMask: true,
  };

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
