import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { FeaturesModule } from './features/features.module';
import { CardModule } from 'primeng/card';
import { CustomToastsModule } from 'shared/components';
import { CoreModule } from './core/core.module';
import { CustomDialogsComponent } from 'shared/components/custom_dialogs/custom-dialogs.component';

@Component({
  standalone: true,
  imports: [
    RouterModule,
    FeaturesModule,
    CardModule,
    CustomToastsModule,
    CoreModule,
    CustomDialogsComponent,
  ],

  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {
  title = 'self-monorepo-project';
}
