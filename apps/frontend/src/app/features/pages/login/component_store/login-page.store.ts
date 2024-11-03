import { Injectable } from '@angular/core';
import { ComponentStore } from '@ngrx/component-store';
import { tapResponse } from '@ngrx/operators';
import { catchError, concatMap, EMPTY, Observable, take, tap } from 'rxjs';
import { UserLoginModel, UserRegisterModel } from 'shared/models';
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

  constructor(private readonly userService: UserService) {
    super({ isLoading: false });
  }

  readonly loginUser = this.effect((trigger$: Observable<UserLoginModel>) => {
    return trigger$.pipe(
      concatMap((userModel) => {
        this.setIsLoading(true);
        return this.userService.loginUser(userModel);
      }),
      tapResponse({
        next: (response) => {
          this.setIsLoading(false);
        },
        error: (error) => {
          this.setIsLoading(false);
        },
      })
    );
  });
}
