import { NgModule } from '@angular/core';
import { PagesModule } from './pages/pages.module';

@NgModule({
  imports: [PagesModule],
  exports: [PagesModule],
})
export class FeaturesModule {}
