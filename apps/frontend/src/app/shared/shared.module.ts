import { NgModule } from '@angular/core';
import { FormModule } from './components/form/form.module';
import { ToastModule } from 'primeng/toast';
import { CustomDialogsModule } from './components';

@NgModule({
  imports: [FormModule, ToastModule, CustomDialogsModule],
  exports: [FormModule, ToastModule, CustomDialogsModule],
})
export class SharedModule {}
