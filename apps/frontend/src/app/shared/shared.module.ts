import { NgModule } from '@angular/core';
import { FormModule } from './components/form/form.module';
import { ToastModule } from 'primeng/toast';

@NgModule({
  imports: [FormModule, ToastModule],
  exports: [FormModule, ToastModule],
})
export class SharedModule {}
