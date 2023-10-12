import { createFeature, createReducer, on } from '@ngrx/store';

import { LocationCity } from '@app/common/models/location.model';

import { LocationActions } from './location.actions';

interface State {
    loading: boolean;
    loaded: boolean;
    cities: LocationCity[];
    errorMsg: string | null;
}

const initialState: State = {
    loading: false,
    loaded: false,
    cities: [],
    errorMsg: null
};

export const locationFeature = createFeature({
    name: 'location',
    reducer: createReducer(
        initialState,
        on(
            LocationActions.load,
            (state): State => ({
                ...state,
                loading: true,
                loaded: false,
                errorMsg: null
            })
        ),
        on(
            LocationActions.loadSuccess,
            (state, { cities }): State => ({
                ...state,
                loading: false,
                loaded: true,
                cities
            })
        ),
        on(
            LocationActions.loadError,
            (state, { errorMsg }): State => ({
                ...state,
                loading: false,
                errorMsg
            })
        )
    )
});
