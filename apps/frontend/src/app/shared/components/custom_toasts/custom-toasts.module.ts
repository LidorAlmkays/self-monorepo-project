import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { BasicToastComponent } from './components/basic-toast/basic-toast.component';
import { ToastModule } from 'primeng/toast';
import { CustomToastsComponent } from './custom-toasts.component';
@NgModule({
  declarations: [CustomToastsComponent, BasicToastComponent],
  imports: [CommonModule, ToastModule],
  exports: [CustomToastsComponent],
})
export class CustomToastsModule {}
