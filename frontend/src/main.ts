import { provideHttpClient } from '@angular/common/http';
import { bootstrapApplication } from '@angular/platform-browser';
import { provideRouter } from '@angular/router';

import { provideStore, provideState } from '@ngrx/store';

import routes from '@app/app-routes';
import { AppComponent } from '@app/app.component';
import { authFeature } from '@app/common/stores/auth/auth.reducer';

bootstrapApplication(AppComponent, {
    providers: [provideHttpClient(), provideRouter(routes), provideStore(), provideState(authFeature)]
}).catch((err) => console.error(err));
