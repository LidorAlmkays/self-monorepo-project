import { Injectable } from '@angular/core';
import {
  CanDeactivate,
  GuardResult,
  MaybeAsync,
  Router,
} from '@angular/router';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root',
})
export class AuthenticateUserTokenGuard implements CanDeactivate<unknown> {
  constructor(
    private readonly router: Router,
    private readonly cookieService: CookieService
  ) {}

  canDeactivate(): MaybeAsync<GuardResult> {
    const session_id = this.cookieService.get('session_id'); // Replace 'token' with your cookie name

    console.log('Found token:' + session_id);
    if (!session_id) {
      this.router.navigate(['/login']); // Redirect to login or a fallback route
      return false; // Deny navigation
    }
    return true;
  }
}
