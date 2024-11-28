import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanDeactivate,
  GuardResult,
  MaybeAsync,
  RouterStateSnapshot,
} from '@angular/router';
import { Observable, of } from 'rxjs';
import { IFormEmptyDataSafe } from 'shared/interfaces';

@Injectable({
  providedIn: 'root',
})
export class FormDiartyGuard implements CanDeactivate<IFormEmptyDataSafe> {
  constructor() {}

  canDeactivate(
    component: IFormEmptyDataSafe,
    currentRoute: ActivatedRouteSnapshot,
    currentState: RouterStateSnapshot,
    nextState: RouterStateSnapshot
  ): MaybeAsync<GuardResult> {
    console.log(!component.isFormEmpty());
    if (!component.isFormEmpty()) {
      //TODO:(lidor) change this to not stop the user but open a dialog that asks the user if he wants to move route and lose his data
      return of(false);
    }
    return of(true);
  }
}
