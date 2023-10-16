import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Company } from '@app/common/models/companies.model';

@Injectable({
    providedIn: 'root'
})
export class AdminCompaniesService {
    constructor(private readonly httpClient: HttpClient) {}

    getCompanies(): Observable<Company[]> {
        // TODO: replace by API call.
        return of([
            { id: 1, name: 'Company 1' },
            { id: 2, name: 'Company 2' }
        ]);
    }
}
