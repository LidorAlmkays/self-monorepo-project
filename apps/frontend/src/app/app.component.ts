import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { FeaturesModule } from './features/features.module';
import { CardModule } from 'primeng/card';

@Component({
  standalone: true,
  imports: [RouterModule, FeaturesModule, CardModule],

  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {
  title = 'self-monorepo-project';
}
