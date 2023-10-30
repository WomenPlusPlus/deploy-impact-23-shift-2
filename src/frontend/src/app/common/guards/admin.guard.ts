import { filter, map, withLatestFrom } from 'rxjs';

import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivateFn, Router } from '@angular/router';

import { Store } from '@ngrx/store';

import { UserKindEnum } from '@app/common/models/users.model';
import { selectProfile } from '@app/common/stores/auth/auth.reducer';
import { isAuthenticated$ } from '@app/common/utils/auth.util';

export const isAdminGuard: CanActivateFn = () => {
    const store = inject(Store);
    const router = inject(Router);
    return isAuthenticated$(store).pipe(
        filter(Boolean),
        withLatestFrom(store.select(selectProfile)),
        map(([, profile]) => profile?.kind),
        map((kind) => {
            if (kind == UserKindEnum.ADMIN) {
                return true;
            }
            return router.parseUrl('/forbidden');
        })
    );
};

export const hasUserIDGuard: CanActivateFn = (route: ActivatedRouteSnapshot) => {
    const store = inject(Store);
    const router = inject(Router);
    const userID = route.params['id'];
    return isAuthenticated$(store).pipe(
        filter(Boolean),
        withLatestFrom(store.select(selectProfile)),
        map(([, profile]) => profile),
        map((profile) => {
            if (profile?.id == userID || profile?.kind == UserKindEnum.ADMIN) {
                return true;
            }
            return router.parseUrl('/forbidden');
        })
    );
};
