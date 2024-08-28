import {
  AfterViewInit,
  ChangeDetectorRef,
  Component,
  EventEmitter,
  Input,
  Output,
} from '@angular/core';
import { FormGroup } from '@angular/forms';
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
  @Output() onSubmit = new EventEmitter();

  public get InputFieldTextTypes(): typeof InputFieldTextTypes {
    return InputFieldTextTypes;
  }
  constructor(private readonly cdr: ChangeDetectorRef) {}
  ngAfterViewInit(): void {
    this.cdr.detectChanges();
  }

  submit() {
    this.onSubmit.emit(this.form.value);
  }
}
