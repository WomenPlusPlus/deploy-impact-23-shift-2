import { createFeature, createReducer, createSelector, on } from '@ngrx/store';

import { AuthActions } from './auth.actions';
import { Profile } from '@app/common/models/profile.model';

interface State {
    loading: boolean;
    account: Profile | null;
    loggedIn: boolean;
}

const initialState: State = {
    loading: true,
    account: null,
    loggedIn: false
};

export const authFeature = createFeature({
    name: 'auth',
    reducer: createReducer(
        initialState,
        on(
            AuthActions.initAuthenticated,
            (state): State => ({ ...state, loggedIn: true })
        ),
        on(
            AuthActions.initNotAuthenticated,
            (state): State => ({ ...state, loading: false, loggedIn: false })
        ),
        on(
            AuthActions.accountLoadedSuccess,
            (state, { account }): State => ({ ...state, loading: false, account })
        ),
        on(
            AuthActions.accountLoadedError,
            (state): State => ({ ...state, loading: false, loggedIn: false })
        )
    ),
    extraSelectors: ({ selectAuthState }) => ({
        selectAuthenticated: createSelector(
            selectAuthState,
            (state) => state.loggedIn && !!state.account,
        )
    })
});

export const { selectLoading: selectAuthLoading, selectAuthenticated, selectAccount: selectProfile } = authFeature;
