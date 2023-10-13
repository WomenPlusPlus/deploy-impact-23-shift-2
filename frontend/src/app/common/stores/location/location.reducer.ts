import { createFeature, createReducer, on } from '@ngrx/store';

import { Language, LocationCity } from '@app/common/models/location.model';

import { LocationActions } from './location.actions';

interface State {
    loadingCities: boolean;
    loadedCities: boolean;
    cities: LocationCity[];
    errorMsgCities: string | null;
    loadingLanguages: boolean;
    loadedLanguages: boolean;
    languages: Language[];
    errorMsgLanguages: string | null;
}

const initialState: State = {
    loadingCities: false,
    loadedCities: false,
    cities: [],
    errorMsgCities: null,
    loadingLanguages: false,
    loadedLanguages: false,
    languages: [],
    errorMsgLanguages: null
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
        ),
        on(
            LocationActions.loadLanguages,
            (state): State => ({
                ...state,
                loadingLanguages: true,
                loadedLanguages: false,
                errorMsgLanguages: null
            })
        ),
        on(
            LocationActions.loadLanguagesSuccess,
            (state, { languages }): State => ({
                ...state,
                loadingLanguages: false,
                loadedLanguages: true,
                languages
            })
        ),
        on(
            LocationActions.loadLanguagesError,
            (state, { errorMsg }): State => ({
                ...state,
                loadingLanguages: false,
                errorMsgLanguages: errorMsg
            })
        )
    )
});

export const { selectCities: selectLocationCities, selectLanguages } = locationFeature;
