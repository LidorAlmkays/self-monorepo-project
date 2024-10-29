import { Injectable } from '@angular/core';
import { ComponentStore } from '@ngrx/component-store';
import { tapResponse } from '@ngrx/operators';
import { catchError, concatMap, EMPTY, Observable, take, tap } from 'rxjs';
import { UserModel } from 'shared/models';
import { UserService } from 'shared/services/user.services';

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

  constructor(private readonly userService: UserService) {
    super({ isLoading: false });
  }

  readonly registerUser = this.effect((trigger$: Observable<UserModel>) => {
    return trigger$.pipe(
      concatMap((userModel) => {
        this.setIsLoading(true);
        return this.userService.registerUser(userModel);
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
