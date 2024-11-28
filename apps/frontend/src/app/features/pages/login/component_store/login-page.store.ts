import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ComponentStore } from '@ngrx/component-store';
import { tapResponse } from '@ngrx/operators';
import { AppPaths, appRoutes } from 'apps/frontend/src/app/app.routes';
import { MessageService } from 'primeng/api';
import { Observable, switchMap } from 'rxjs';
import { CustomToastsKeys } from 'shared/components';
import { SeverityTypes } from 'shared/components/custom_toasts/enums/severity-types.enum';
import { UserLoginModel } from 'shared/models';
import { UserService } from 'shared/services/user.services';

export interface LoginState {
  isLoading: boolean;
}

@Injectable()
export class LoginStore extends ComponentStore<LoginState> {
  private readonly isLoading$: Observable<boolean> = this.select(
    (state) => state.isLoading
  );
  readonly vm$ = this.select({
    isLoading: this.isLoading$,
  });

  private readonly setIsLoading = this.updater((state, isLoading: boolean) => {
    const newState: LoginState = {
      ...state,
      isLoading,
    };
    return newState;
  });

  constructor(
    private readonly userService: UserService,
    private readonly router: Router,
    private readonly messageService: MessageService
  ) {
    super({ isLoading: false });
  }

  readonly loginUser = this.effect((trigger$: Observable<UserLoginModel>) => {
    return trigger$.pipe(
      switchMap((userModel) => {
        this.setIsLoading(true);
        return this.userService.loginUser(userModel).pipe(
          tapResponse(
            (response) => {
              this.setIsLoading(false);
              this.router.navigate(['/' + AppPaths.home()]);
              this.messageService.add({
                key: CustomToastsKeys.BasicToast,
                severity: SeverityTypes.SUCCESS,
                summary: 'Login Successfully',
                detail: 'User successfully logged in.',
              });
            },
            (error) => {
              this.setIsLoading(false);
              this.messageService.add({
                key: CustomToastsKeys.BasicToast,
                severity: SeverityTypes.ERROR,
                summary: 'Login Failed',
                detail: 'User failed to login in.',
              });
            }
          )
        );
      })
    );
  });
}
