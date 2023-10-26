import { iif, map, Observable, of, switchMap } from 'rxjs';

import { Pipe, PipeTransform } from '@angular/core';

import { Store } from '@ngrx/store';

import { Profile } from '@app/common/models/profile.model';
import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';
import { selectProfile } from '@app/common/stores/auth/auth.reducer';
import { isAuthenticated$ } from '@app/common/utils/auth.util';

@Pipe({
    name: 'isAuthorized',
    standalone: true
})
export class IsAuthorizedPipe implements PipeTransform {
    constructor(private readonly store: Store) {}

    transform(
        isTrue: boolean,
        kinds: UserKindEnum | UserKindEnum[] | null = null,
        roles: UserRoleEnum | UserRoleEnum[] | null = null
    ): Observable<boolean> {
        return isAuthenticated$(this.store).pipe(
            switchMap((authenticated) =>
                iif(
                    () => authenticated,
                    this.store
                        .select(selectProfile)
                        .pipe(map((profile) => this.isAuthorized(profile, isTrue, kinds, roles))),
                    of(!isTrue)
                )
            )
        );
    }

    private isAuthorized(
        profile: Profile | null,
        isTrue: boolean,
        kind: UserKindEnum | UserKindEnum[] | null,
        role: UserRoleEnum | UserRoleEnum[] | null
    ): boolean {
        if (!profile) {
            return !isTrue;
        }
        if (kind) {
            const kinds = Array.isArray(kind) ? kind : [kind];
            const found = kinds.includes(profile.kind);

            if (found !== isTrue) {
                return false;
            }
        }
        if (role) {
            if (!profile.role) {
                return !isTrue;
            }

            const roles = Array.isArray(role) ? role : [role];
            const found = roles.includes(profile.role);

            if (found !== isTrue) {
                return false;
            }
        }
        return true;
    }
}
