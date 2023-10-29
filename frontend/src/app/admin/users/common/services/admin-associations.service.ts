import { map, Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { Association } from '@app/common/models/associations.model';
import { AssociationList } from '@app/dashboard/common/models/associations.model';

@Injectable({
    providedIn: 'root'
})
export class AdminAssociationsService {
    constructor(private readonly httpClient: HttpClient) {}

    getAssociations(): Observable<Association[]> {
        return this.httpClient
            .get<AssociationList>(`${environment.API_BASE_URL}/api/v1/associations`)
            .pipe(map(({ items }) => items));
    }
}
