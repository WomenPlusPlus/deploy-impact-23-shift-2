import { CanActivateFn, Router } from '@angular/router';
import { inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { map } from 'rxjs';
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
                return router.parseUrl('/admin');
            }
            return true;
        })
    );
};
