import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeStore } from './component_store/home-page.store';
import { CardModule } from 'primeng/card';

@Component({
  selector: 'app-home-page',
  standalone: true,
  providers: [HomeStore],
  imports: [CommonModule, CardModule],

  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.css',
})
export class HomePageComponent {}
