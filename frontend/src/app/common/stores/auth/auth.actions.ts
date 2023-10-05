import { createActionGroup, emptyProps } from '@ngrx/store';

export const AuthActions = createActionGroup({
    source: 'Auth Page',
    events: {
        Login: emptyProps()
    }
});
