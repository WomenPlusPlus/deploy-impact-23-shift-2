import { createFeature, createReducer, on } from '@ngrx/store';

import { LocationCity } from '@app/common/models/location.model';

import { LocationActions } from './location.actions';

interface State {
    loadingCities: boolean;
    loadedCities: boolean;
    cities: LocationCity[];
    errorMsgCities: string | null;
}

const initialState: State = {
    loadingCities: false,
    loadedCities: false,
    cities: [],
    errorMsgCities: null
};

export const locationFeature = createFeature({
    name: 'location',
    reducer: createReducer(
        initialState,
        on(
            LocationActions.loadCities,
            (state): State => ({
                ...state,
                loadingCities: true,
                loadedCities: false,
                errorMsgCities: null
            })
        ),
        on(
            LocationActions.loadCitiesSuccess,
            (state, { cities }): State => ({
                ...state,
                loadingCities: false,
                loadedCities: true,
                cities
            })
        ),
        on(
            LocationActions.loadCitiesError,
            (state, { errorMsg }): State => ({
                ...state,
                loadingCities: false,
                errorMsgCities: errorMsg
            })
        )
    )
});

export const { selectCities: selectLocationCities } = locationFeature;
