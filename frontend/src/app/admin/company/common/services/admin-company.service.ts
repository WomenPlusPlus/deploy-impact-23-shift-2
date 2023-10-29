import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { CompanyProfileModel } from '@app/companies/profile/common/models/company-profile.model';

@Injectable({
    providedIn: 'root'
})
export class AdminCompanyService {
    constructor(private readonly httpClient: HttpClient) {}

    createCompany(payload: FormData): Observable<{ id: number }> {
        // post company data to back
        return this.httpClient.post<{ id: number }>(`${environment.API_BASE_URL}/api/v1/companies`, payload);
    }

    editCompany(payload: FormData, id: number): Observable<void> {
        // update company data to back end
        return this.httpClient.put<void>(`${environment.API_BASE_URL}/api/v1/companies/${id}`, payload);
    }

    getCompany(id: number): Observable<CompanyProfileModel> {
        return this.httpClient.get<CompanyProfileModel>(`${environment.API_BASE_URL}/api/v1/companies/${id}`);
    }
}
