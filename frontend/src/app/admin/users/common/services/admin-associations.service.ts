import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Association } from '@app/common/models/associations.model';

@Injectable({
    providedIn: 'root'
})
export class AdminAssociationsService {
    constructor(private readonly httpClient: HttpClient) {}

    getAssociations(): Observable<Association[]> {
        // TODO: replace by API call.
        return of([
            { id: 1, name: 'Association 1' },
            { id: 2, name: 'Association 2' }
        ]);
    }
}
