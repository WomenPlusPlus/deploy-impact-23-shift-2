import { provideHotToastConfig } from '@ngneat/hot-toast';

import { provideHttpClient } from '@angular/common/http';
import { bootstrapApplication } from '@angular/platform-browser';
import { provideRouter } from '@angular/router';

import { provideEffects } from '@ngrx/effects';
import { provideState, provideStore } from '@ngrx/store';

import routes from '@app/app-routes';
import { AppComponent } from '@app/app.component';
import { authFeature } from '@app/common/stores/auth/auth.reducer';
import * as LocationEffects from '@app/common/stores/location/location.effects';
import { locationFeature } from '@app/common/stores/location/location.reducer';

bootstrapApplication(AppComponent, {
    providers: [
        provideHttpClient(),
        provideRouter(routes),
        provideStore(),
        provideState(authFeature),
        provideState(locationFeature),
        provideEffects(LocationEffects),
        provideHotToastConfig()
    ]
}).catch((err) => console.error(err));
