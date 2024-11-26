import { Injectable } from '@angular/core';
import { ComponentStore } from '@ngrx/component-store';

export interface HomeState {}

@Injectable()
export class HomeStore extends ComponentStore<HomeState> {
  readonly vm$ = this.select({});

  constructor() {
    super({});
  }
}
