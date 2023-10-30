import Fuse from 'fuse.js';
import { EMPTY, map, Observable, startWith } from 'rxjs';

import { Pipe, PipeTransform } from '@angular/core';

import { Language } from '@app/common/models/location.model';

import FuseResult = Fuse.FuseResult;

@Pipe({
    name: 'filterLanguage',
    standalone: true
})
export class FilterLanguagePipe implements PipeTransform {
    transform(languages: Language[] | null, term$: Observable<string | null>): Observable<FuseResult<Language>[]> {
        if (!languages) {
            return EMPTY;
        }

        const fuse = new Fuse(languages, {
            keys: ['name', 'shortName'],
            shouldSort: false,
            threshold: 0.2
        });
        return term$.pipe(
            startWith(''),
            map((term) =>
                !term || term.length < 3
                    ? languages.map((item, refIndex) => ({ refIndex, item }))
                    : fuse.search(term || '.')
            )
        );
    }
}
