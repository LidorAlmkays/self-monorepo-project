import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CUstomDialogsKeys as CustomDialogsKeys } from './custom-dialogs-keys.enum';
import { DialogModule } from 'primeng/dialog';
import { ConfirmDialogModule } from 'primeng/confirmdialog';

@Component({
  selector: 'app-custom-dialogs',
  standalone: true,
  imports: [CommonModule, DialogModule, ConfirmDialogModule],
  templateUrl: './custom-dialogs.component.html',
  styleUrl: './custom-dialogs.component.css',
})
export class CustomDialogsComponent {
  basicConfirmDialogKey = CustomDialogsKeys.BasicConfirmDialog;
}
