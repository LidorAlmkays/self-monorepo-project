import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'apps/frontend/src/environments/environment';
import {
  catchError,
  first,
  map,
  Observable,
  of,
  switchMap,
  tap,
  throwError,
} from 'rxjs';
import { UserLoginModel, UserRegisterModel } from 'shared/models';

@Injectable()
export class UserService {
  constructor(private readonly http: HttpClient) {}
  private readonly userGatewayUrl = environment.gateway + '/user';
  loginUser(user: UserLoginModel): Observable<void> {
    return this.http
      .post(this.userGatewayUrl + '/login', user, {
        headers: {
          'Content-Type': 'application/json',
        },
      })
      .pipe(
        first(),
        tap((cookie) => {
          console.log('Recived cookie: ' + cookie);
          console.log('User logged in.');
        }),
        map(() => {}), // Map the successful response to `void`
        catchError((error) => {
          console.error('Error registering user:', error);
          return throwError(() => error);
        })
      );
  }

  registerUser(user: UserRegisterModel): Observable<UserRegisterModel> {
    return this.http
      .post(this.userGatewayUrl + '/register', user, {
        headers: {
          'Content-Type': 'application/json',
        },
      })
      .pipe(
        first(),
        tap(() => {
          console.log('User registered successfully.');
        }),
        switchMap(() => of(user)), // Use `switchMap` to return `username` from `user`
        catchError((error) => {
          console.error('Error registering user:', error);
          return throwError(() => error);
        })
      );
  }
}
