import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UserModel } from 'shared/models';
import { Observable, of } from 'rxjs';
@Injectable()
export class UserService {
  constructor(private http: HttpClient) {}

  registerUser(userModel: UserModel): Observable<UserModel> {
    return of(userModel);
  }
}
