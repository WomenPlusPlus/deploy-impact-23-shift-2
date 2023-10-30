import { filter, map, withLatestFrom } from 'rxjs';

import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';

import { Store } from '@ngrx/store';

import { UserStateEnum } from '@app/common/models/users.model';
import { selectProfile } from '@app/common/stores/auth/auth.reducer';
import { isAuthenticated$ } from '@app/common/utils/auth.util';

export const activatedGuard: CanActivateFn = () => {
    const store = inject(Store);
    const router = inject(Router);
    return isAuthenticated$(store).pipe(
        filter(Boolean),
        withLatestFrom(store.select(selectProfile)),
        map(([, profile]) => profile?.state || UserStateEnum.DELETED),
        map((state) => {
            switch (state) {
                case UserStateEnum.ACTIVE:
                    return true;
                case UserStateEnum.INVITED:
                    return router.parseUrl('/setup');
                default:
                    return false;
            }
        })
    );
};

export const invitedGuard: CanActivateFn = () => {
    const store = inject(Store);
    const router = inject(Router);
    return isAuthenticated$(store).pipe(
        filter(Boolean),
        withLatestFrom(store.select(selectProfile)),
        map(([, profile]) => profile?.state || UserStateEnum.DELETED),
        map((state) => {
            switch (state) {
                case UserStateEnum.ACTIVE:
                    return router.parseUrl('/');
                case UserStateEnum.INVITED:
                    return true;
                default:
                    return false;
            }
        })
    );
};
