import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ComponentStore } from '@ngrx/component-store';
import { tapResponse } from '@ngrx/operators';
import { AppPaths } from 'apps/frontend/src/app/app.routes';
import { catchError, concatMap, EMPTY, Observable, take, tap } from 'rxjs';
import { UserRegisterModel } from 'shared/models';
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

  constructor(
    private readonly userService: UserService,
    private router: Router
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
                this.userService.loginUser(user);
                this.setIsLoading(false);
                this.router.navigate(['/' + AppPaths.home()]);
              },
              error: (error) => {
                //TODO:(lidor) add an error with toast why failed
                this.setIsLoading(false);
              },
            })
          );
        })
      );
    }
  );
}
