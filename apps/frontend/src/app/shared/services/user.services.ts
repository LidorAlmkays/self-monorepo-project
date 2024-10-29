import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'apps/frontend/src/environments/environment';
import { UserModel } from 'shared/models';

@Injectable()
export class UserService {
  constructor(private readonly http: HttpClient) {}
  async registerUser(user: UserModel) {
    try {
      this.http
        .put(environment.gateway + '/user', user, {
          headers: {
            'Content-Type': 'application/json',
          },
        })
        .subscribe((response) => {
          console.log(response);
        });
    } catch (error) {
      console.error('Error fetching quotes:', error);
    }
  }
}
