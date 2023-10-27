import { AuthService } from '@auth0/auth0-angular';
import { HotToastService } from '@ngneat/hot-toast';
import { catchError, EMPTY, exhaustMap, first, map, of, switchMap, tap } from 'rxjs';

import { inject } from '@angular/core';
import { Router } from '@angular/router';

import { Actions, createEffect, ofType, ROOT_EFFECTS_INIT } from '@ngrx/effects';

import { AccountService } from '@app/common/services/account.service';
import { AuthActions } from '@app/common/stores/auth/auth.actions';

const loadAuthentication = createEffect(
    (actions$ = inject(Actions), authService = inject(AuthService)) => {
        return actions$.pipe(
            ofType(ROOT_EFFECTS_INIT),
            switchMap(() => authService.isAuthenticated$.pipe(first())),
            map((authenticated) =>
                authenticated ? AuthActions.initAuthenticated() : AuthActions.initNotAuthenticated()
            )
        );
    },
    { functional: true }
);

const initAuthenticated = createEffect(
    (
        actions$ = inject(Actions),
        router = inject(Router),
        toast = inject(HotToastService),
        accountService = inject(AccountService)
    ) => {
        return actions$.pipe(
            ofType(AuthActions.initAuthenticated),
            switchMap(() =>
                accountService.me().pipe(
                    map((account) => AuthActions.accountLoadedSuccess({ account })),
                    catchError((error) => {
                        console.error(error);
                        toast.error(
                            "Could not load account details, you were logged off from the application! Please try again later or contact the site's administrator."
                        );
                        return of(AuthActions.accountLoadedError()).pipe(tap(() => router.navigate(['/login'])));
                    })
                )
            )
        );
    },
    { functional: true }
);

const initNotAuthenticated = createEffect(
    (actions$ = inject(Actions), authService = inject(AuthService)) => {
        return actions$.pipe(
            ofType(AuthActions.initNotAuthenticated),
            exhaustMap(() => authService.isAuthenticated$),
            switchMap((authenticated) => (authenticated ? authService.logout() : EMPTY))
        );
    },
    { functional: true, dispatch: false }
);

export { loadAuthentication, initAuthenticated, initNotAuthenticated };
