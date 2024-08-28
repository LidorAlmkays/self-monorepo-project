import { Injectable } from '@angular/core';
import { ComponentStore } from '@ngrx/component-store';
import { UserModel } from 'shared/models';

export interface RegisterState {
  userData?: UserModel;
}

@Injectable()
export class RegisterStore extends ComponentStore<RegisterState> {
  constructor() {
    super();
  }
}
