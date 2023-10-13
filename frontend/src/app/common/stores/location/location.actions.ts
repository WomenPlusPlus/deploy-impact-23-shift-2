import { createActionGroup, emptyProps, props } from '@ngrx/store';

import { Language, LocationCity } from '@app/common/models/location.model';

export const LocationActions = createActionGroup({
    source: 'Location',
    events: {
        loadInitialData: emptyProps(),
        LoadCities: emptyProps(),
        loadCitiesSuccess: props<{ cities: LocationCity[] }>(),
        loadCitiesError: props<{ errorMsg: string }>(),
        LoadLanguages: emptyProps(),
        loadLanguagesSuccess: props<{ languages: Language[] }>(),
        loadLanguagesError: props<{ errorMsg: string }>()
    }
});
