import { createActionGroup, emptyProps, props } from '@ngrx/store';

import { LocationCity } from '@app/common/models/location.model';

export const LocationActions = createActionGroup({
    source: 'Location',
    events: {
        LoadCities: emptyProps(),
        loadCitiesSuccess: props<{ cities: LocationCity[] }>(),
        loadCitiesError: props<{ errorMsg: string }>()
    }
});
