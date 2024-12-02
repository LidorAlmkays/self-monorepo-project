import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ComponentStore } from '@ngrx/component-store';
import { tapResponse } from '@ngrx/operators';
import { Store } from '@ngrx/store';
import { MessageService } from 'primeng/api';
import { concatMap, Observable } from 'rxjs';
import { CustomToastsKeys } from 'shared/components';
import { SeverityTypes } from 'shared/components/custom_toasts/enums/severity-types.enum';
import { UserRegisterModel } from 'shared/models';
import { UserService } from 'shared/services/user.services';
import { userRegisterActions } from '../../../ngrx_store/user/actions/user-register.action';

export interface RegisterState {
  isLoading: boolean;
}

@Injectable()
export class RegisterStore extends ComponentStore<RegisterState> {
  private readonly isLoading$: Observable<boolean> = this.select(
    (state) => state.isLoading
  );
  readonly vm$ = this.select({
    isLoading: this.isLoading$,
  });

  private readonly setIsLoading = this.updater((state, isLoading: boolean) => {
    const newState: RegisterState = {
      ...state,
      isLoading,
    };
    return newState;
  });

  constructor(
    private readonly store: Store,
    private readonly userService: UserService,
    private router: Router,
    private readonly messageService: MessageService
  ) {
    super({ isLoading: false });
  }

  readonly registerUser = this.effect(
    (trigger$: Observable<UserRegisterModel>) => {
      return trigger$.pipe(
        concatMap((user) => {
          this.setIsLoading(true);
          return this.userService.registerUser(user).pipe(
            tapResponse({
              next: (user) => {
                this.userService.loginUser(user).forEach(() => {
                  this.store.dispatch(userRegisterActions.loggedIn());
                });
                this.setIsLoading(false);
                this.messageService.add({
                  key: CustomToastsKeys.BasicToast,
                  severity: SeverityTypes.SUCCESS,
                  summary: 'Register Successfully',
                  detail: 'User successfully register in.',
                });
              },
              error: (error) => {
                this.messageService.add({
                  key: CustomToastsKeys.BasicToast,
                  severity: SeverityTypes.ERROR,
                  summary: 'Register Failed',
                  detail: 'User failed to register.',
                });
                this.setIsLoading(false);
              },
            })
          );
        })
      );
    }
  );
}
