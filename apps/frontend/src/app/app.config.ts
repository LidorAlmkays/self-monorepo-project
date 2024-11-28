import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { appRoutes } from './app.routes';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { provideStore } from '@ngrx/store';
import { provideStoreDevtools } from '@ngrx/store-devtools';
import { environment } from '../environments/environment.production';
import { provideRouterStore, routerReducer } from '@ngrx/router-store';
import { CustomSerializer } from './features/ngrx_store/router_serializers/custom-route-serializer';
import { CredentialsInterceptor } from './core/intersepters/credentials-interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    provideAnimations(),
    provideHttpClient(withInterceptors([CredentialsInterceptor])),
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(appRoutes),
    provideStore({ router: routerReducer }),
    provideRouterStore({ serializer: CustomSerializer }),
    provideStoreDevtools({
      maxAge: 25, // Retains last 25 states
      logOnly: !environment.production, // Restrict extension to log-only mode
      autoPause: true, // Pauses recording actions and state changes when the extension window is not open
      trace: true, //  If set to true, will include stack trace for every dispatched action, so you can see it in trace tab jumping directly to that part of code
      traceLimit: 75, // maximum stack trace frames to be stored (in case trace option was provided as true)
      connectInZone: true, // If set to true, the connection is established within the Angular zone
    }),
  ],
};
