import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CustomToastsKeys } from '../../enums';
import { IBaseToastRequirements } from '../../base-toast-requirements';

@Component({
  selector: 'app-basic-toast',
  templateUrl: './basic-toast.component.html',
  styleUrl: './basic-toast.component.css',
})
export class BasicToastComponent implements IBaseToastRequirements {
  public key: CustomToastsKeys = CustomToastsKeys.BasicToast;
}
