import { HttpEvent, HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';

import { first, Observable } from 'rxjs';
import { mergeMap, switchMap } from 'rxjs/operators';
import { AuthService } from '@auth0/auth0-angular';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
    const auth = inject(AuthService);

    if (checkPublicRoute(req.url)) {
        return next(req);
    }

    const auth0Interceptor = (): Observable<HttpEvent<unknown>> => {
        return auth.idTokenClaims$.pipe(
            mergeMap((token) =>
                next(
                    req.clone({
                        setHeaders: {
                            Authorization: `Bearer ${token?.__raw}`
                        }
                    })
                )
            )
        );
    };


    return auth.isAuthenticated$.pipe(
        first(),
        switchMap((authenticated) => !authenticated
            ? next(req)
            : auth0Interceptor()
        )
    );
};

const PUBLIC_ROUTES: string[] = ['https://shift2-deployimpact.eu.auth0.com/', '/locations'];

function checkPublicRoute(url: string) {
    return PUBLIC_ROUTES.some((path) => url.includes(path));
}
