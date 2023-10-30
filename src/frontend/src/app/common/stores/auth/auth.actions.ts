import { createActionGroup, emptyProps, props } from '@ngrx/store';

import { Profile } from '@app/common/models/profile.model';

export const AuthActions = createActionGroup({
    source: 'Auth Page',
    events: {
        InitAuthenticated: emptyProps(),
        InitNotAuthenticated: emptyProps(),
        AccountLoadedSuccess: props<{ account: Profile }>(),
        AccountLoadedError: emptyProps()
    }
});
