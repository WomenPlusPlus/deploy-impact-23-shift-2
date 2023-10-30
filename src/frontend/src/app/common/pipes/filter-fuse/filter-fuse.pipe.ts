import Fuse from 'fuse.js';
import { EMPTY, map, Observable, startWith } from 'rxjs';

import { Pipe, PipeTransform } from '@angular/core';

import FuseResult = Fuse.FuseResult;

@Pipe({
    name: 'filterFuse',
    standalone: true
})
export class FilterFusePipe<T> implements PipeTransform {
    transform(value: T[] | null, term$: Observable<string | null>, keys: string[]): Observable<FuseResult<T>[]> {
        if (!value) {
            return EMPTY;
        }

        const fuse = new Fuse(value, {
            keys,
            shouldSort: false,
            threshold: 0.2
        });
        return term$.pipe(
            startWith(''),
            map((term) =>
                !term || term.length < 3
                    ? value.map((item, refIndex) => ({ refIndex, item }))
                    : fuse.search(term || '.')
            )
        );
    }
}
