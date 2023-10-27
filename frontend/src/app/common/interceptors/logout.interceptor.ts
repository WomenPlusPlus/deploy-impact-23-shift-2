import { AuthService } from '@auth0/auth0-angular';
import { throwError } from 'rxjs';
import { catchError, switchMap } from 'rxjs/operators';

import { HttpInterceptorFn, HttpStatusCode } from '@angular/common/http';
import { inject } from '@angular/core';

export const logoutInterceptor: HttpInterceptorFn = (req, next) => {
    const auth = inject(AuthService);

    return next(req).pipe(
        catchError((error) => {
            if (error.status === HttpStatusCode.Unauthorized || req.url.endsWith('/api/v1/me')) {
                auth.logout().pipe(switchMap(() => throwError(() => error)));
            }

            return throwError(() => error);
        })
    );
};
