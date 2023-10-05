import { createFeature, createReducer, on } from '@ngrx/store';

import { AuthActions } from './auth.actions';

interface State {
    loggedIn: boolean;
}

const initialState: State = {
    loggedIn: false
};

export const authFeature = createFeature({
    name: 'auth',
    reducer: createReducer(
        initialState,
        on(
            AuthActions.login,
            (state): State => ({
                ...state,
                loggedIn: !state.loggedIn
            })
        )
    )
});
