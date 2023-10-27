import { exhaustMap, filter, Observable, of, switchMap } from 'rxjs';

import { Store } from '@ngrx/store';

import { Profile } from '@app/common/models/profile.model';
import { selectAuthenticated, selectAuthLoading, selectProfile } from '@app/common/stores/auth/auth.reducer';

export const isAuthenticated$ = (store: Store): Observable<boolean> =>
    store.select(selectAuthLoading).pipe(
        filter((loading) => !loading),
        exhaustMap(() => store.select(selectAuthenticated))
    );

export const profile$ = (store: Store): Observable<Profile | null> =>
    isAuthenticated$(store).pipe(
        switchMap((authenticated) => (authenticated ? store.select(selectProfile) : of(null)))
    );
