import { exhaustMap, filter, Observable } from 'rxjs';
import { Store } from '@ngrx/store';
import { selectAuthenticated, selectAuthLoading } from '@app/common/stores/auth/auth.reducer';

export const isAuthenticated$ = (store: Store): Observable<boolean> => store.select(selectAuthLoading).pipe(
    filter((loading) => !loading),
    exhaustMap(() => store.select(selectAuthenticated))
);
