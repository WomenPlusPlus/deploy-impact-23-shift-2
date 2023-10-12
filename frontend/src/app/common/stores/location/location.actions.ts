import { createActionGroup, emptyProps, props } from '@ngrx/store';

import { LocationCity } from '@app/common/models/location.model';

export const LocationActions = createActionGroup({
    source: 'Locations',
    events: {
        Load: emptyProps(),
        loadSuccess: props<{ cities: LocationCity[] }>(),
        loadError: props<{ errorMsg: string }>()
    }
});
