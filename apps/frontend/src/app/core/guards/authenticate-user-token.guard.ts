import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class AuthenticateUserTokenGuard implements CanActivate {
  constructor(
    private readonly router: Router,
    private readonly cookieService: CookieService
  ) {}

  canActivate(): Observable<boolean> | Promise<boolean> | boolean {
    const session_id = this.cookieService.get('session_id'); // Replace 'session_id' with your cookie name

    console.debug('Found token:' + session_id);
    if (!session_id) {
      this.router.navigate(['/login']); // Redirect to login or a fallback route
      return false; // Deny access
    }
    return true; // Allow access
  }
}
