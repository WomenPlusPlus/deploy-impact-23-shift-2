import { filter, map, withLatestFrom } from 'rxjs';

import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivateFn, Router } from '@angular/router';

import { Store } from '@ngrx/store';

import { UserKindEnum } from '@app/common/models/users.model';
import { selectProfile } from '@app/common/stores/auth/auth.reducer';
import { isAuthenticated$ } from '@app/common/utils/auth.util';

export const isRelatedToCompany: CanActivateFn = (route: ActivatedRouteSnapshot) => {
    const store = inject(Store);
    const router = inject(Router);
    const pageId = route.params['id'];
    return isAuthenticated$(store).pipe(
        filter(Boolean),
        withLatestFrom(store.select(selectProfile)),
        map(([, profile]) => profile),
        map((profile) => {
            if (profile?.companyId == pageId || profile?.kind == UserKindEnum.ADMIN) {
                return true;
            }
            return router.parseUrl('/forbidden');
        })
    );
};

export const isRelatedToAssociation: CanActivateFn = (route: ActivatedRouteSnapshot) => {
    const store = inject(Store);
    const router = inject(Router);
    const pageId = route.params['id'];
    return isAuthenticated$(store).pipe(
        filter(Boolean),
        withLatestFrom(store.select(selectProfile)),
        map(([, profile]) => profile),
        map((profile) => {
            if (profile?.associationId == pageId || profile?.kind == UserKindEnum.ADMIN) {
                return true;
            }
            return router.parseUrl('/forbidden');
        })
    );
};
