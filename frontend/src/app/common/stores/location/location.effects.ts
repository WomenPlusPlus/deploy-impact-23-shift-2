import { catchError, exhaustMap, map, of, switchMap } from 'rxjs';

import { inject } from '@angular/core';

import { Actions, createEffect, ofType } from '@ngrx/effects';

import { LocationsService } from '@app/common/services/locations/locations.service';

import { LocationActions } from './location.actions';

const loadLocationInitialData = createEffect(
    (actions$ = inject(Actions)) => {
        return actions$.pipe(
            ofType(LocationActions.loadInitialData),
            switchMap(() => [LocationActions.loadCities(), LocationActions.loadLanguages()])
        );
    },
    { functional: true }
);

const loadLocationCities = createEffect(
    (actions$ = inject(Actions), locationsService = inject(LocationsService)) => {
        return actions$.pipe(
            ofType(LocationActions.loadCities),
            exhaustMap(() =>
                locationsService.getCities().pipe(
                    map((cities) => LocationActions.loadCitiesSuccess({ cities })),
                    catchError((error: { message: string }) =>
                        of(LocationActions.loadCitiesError({ errorMsg: error.message }))
                    )
                )
            )
        );
    },
    { functional: true }
);

const loadLocationLanguages = createEffect(
    (actions$ = inject(Actions), locationsService = inject(LocationsService)) => {
        return actions$.pipe(
            ofType(LocationActions.loadLanguages),
            exhaustMap(() =>
                locationsService.getLanguages().pipe(
                    map((languages) => LocationActions.loadLanguagesSuccess({ languages })),
                    catchError((error: { message: string }) =>
                        of(LocationActions.loadLanguagesError({ errorMsg: error.message }))
                    )
                )
            )
        );
    },
    { functional: true }
);

export { loadLocationInitialData, loadLocationCities, loadLocationLanguages };
