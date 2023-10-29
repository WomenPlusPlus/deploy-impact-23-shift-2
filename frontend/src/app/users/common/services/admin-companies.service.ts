import { map, Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { Company } from '@app/common/models/companies.model';
import { CompanyList } from '@app/companies/profile/common/models/company-profile.model';

@Injectable({
    providedIn: 'root'
})
export class AdminCompaniesService {
    constructor(private readonly httpClient: HttpClient) {}

    getCompanies(): Observable<Company[]> {
        return this.httpClient
            .get<CompanyList>(`${environment.API_BASE_URL}/api/v1/companies`)
            .pipe(map(({ items }) => items));
    }
}
