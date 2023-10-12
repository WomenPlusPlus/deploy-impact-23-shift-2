import { catchError, exhaustMap, map, of } from 'rxjs';

import { inject } from '@angular/core';

import { Actions, createEffect, ofType } from '@ngrx/effects';

import { LocationsService } from '@app/common/services/locations/locations.service';

import { LocationActions } from './location.actions';

export const loadActors = createEffect(
    (actions$ = inject(Actions), locationsService = inject(LocationsService)) => {
        return actions$.pipe(
            ofType(LocationActions.load),
            exhaustMap(() =>
                locationsService.getCities().pipe(
                    map((cities) => LocationActions.loadSuccess({ cities })),
                    catchError((error: { message: string }) =>
                        of(LocationActions.loadError({ errorMsg: error.message }))
                    )
                )
            )
        );
    },
    { functional: true }
);
