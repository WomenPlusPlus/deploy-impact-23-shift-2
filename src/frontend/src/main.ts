import { provideAuth0 } from '@auth0/auth0-angular';
import { provideHotToastConfig } from '@ngneat/hot-toast';

import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { bootstrapApplication } from '@angular/platform-browser';
import { provideRouter } from '@angular/router';

import { provideEffects } from '@ngrx/effects';
import { provideState, provideStore } from '@ngrx/store';

import environment from '@envs/environment';

import routes from '@app/app-routes';
import { AppComponent } from '@app/app.component';
import { authInterceptor } from '@app/common/interceptors/auth.interceptor';
import { logoutInterceptor } from '@app/common/interceptors/logout.interceptor';
import * as AuthEffects from '@app/common/stores/auth/auth.effects';
import { authFeature } from '@app/common/stores/auth/auth.reducer';
import * as LocationEffects from '@app/common/stores/location/location.effects';
import { locationFeature } from '@app/common/stores/location/location.reducer';

bootstrapApplication(AppComponent, {
    providers: [
        provideHttpClient(withInterceptors([authInterceptor, logoutInterceptor])),
        provideRouter(routes),
        provideStore(),
        provideState(authFeature),
        provideEffects(AuthEffects),
        provideState(locationFeature),
        provideEffects(LocationEffects),
        provideHotToastConfig(),
        provideAuth0({
            domain: environment.AUTH0_DOMAIN,
            clientId: environment.AUTH0_CLIENT_ID,
            authorizationParams: {
                redirect_uri: window.location.origin
            }
        })
    ]
}).catch((err) => console.error(err));
