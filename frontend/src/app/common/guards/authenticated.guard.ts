import { map } from 'rxjs';

import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';

import { Store } from '@ngrx/store';

import { isAuthenticated$ } from '@app/common/utils/auth.util';

export const authenticatedGuard: CanActivateFn = () => {
    const store = inject(Store);
    const router = inject(Router);
    return isAuthenticated$(store).pipe(
        map((authenticated) => {
            if (authenticated) {
                return true;
            }
            return router.parseUrl('/login');
        })
    );
};

export const notAuthenticatedGuard: CanActivateFn = () => {
    const store = inject(Store);
    const router = inject(Router);
    return isAuthenticated$(store).pipe(
        map((authenticated) => {
            if (authenticated) {
                return router.parseUrl('/dashboard');
            }
            return true;
        })
    );
};
