import Fuse from 'fuse.js';
import { EMPTY, map, Observable, startWith } from 'rxjs';

import { Pipe, PipeTransform } from '@angular/core';

import { LocationCity } from '@app/common/models/location.model';

import FuseResult = Fuse.FuseResult;

@Pipe({
    name: 'filterCity',
    standalone: true
})
export class FilterCityPipe implements PipeTransform {
    transform(cities: LocationCity[] | null, term$: Observable<string | null>): Observable<FuseResult<LocationCity>[]> {
        if (!cities) {
            return EMPTY;
        }

        const fuse = new Fuse(cities, {
            keys: ['name'],
            shouldSort: false,
            threshold: 0.2
        });
        return term$.pipe(
            startWith(''),
            map((term) =>
                !term || term.length < 3
                    ? cities.map((item, refIndex) => ({ refIndex, item }))
                    : fuse.search(term || '.')
            )
        );
    }
}
