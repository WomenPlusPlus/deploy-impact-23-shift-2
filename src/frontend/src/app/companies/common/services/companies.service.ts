import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { CompanyList, CompanyItem } from '@app/companies/profile/common/models/company-profile.model';
import { JobList } from '@app/jobs/common/models/job.model';

@Injectable({
    providedIn: 'root'
})
export class CompaniesService {
    constructor(private readonly httpClient: HttpClient) {}

    createCompany(payload: FormData): Observable<{ id: number }> {
        return this.httpClient.post<{ id: number }>(`${environment.API_BASE_URL}/api/v1/companies`, payload);
    }

    editCompany(payload: FormData, id: number): Observable<void> {
        return this.httpClient.put<void>(`${environment.API_BASE_URL}/api/v1/companies/${id}`, payload);
    }

    getCompany(id: number): Observable<CompanyItem> {
        return this.httpClient.get<CompanyItem>(`${environment.API_BASE_URL}/api/v1/companies/${id}`);
    }

    getCompaniesList(): Observable<CompanyList> {
        return this.httpClient.get<CompanyList>(`${environment.API_BASE_URL}/api/v1/companies`);
    }

    deleteCompany(id: number): Observable<void> {
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/companies/${id}`);
    }

    getJobsByCompany(id: number): Observable<JobList> {
        return this.httpClient.get<JobList>(`${environment.API_BASE_URL}/api/v1/companies/${id}/jobs`);
    }
}
