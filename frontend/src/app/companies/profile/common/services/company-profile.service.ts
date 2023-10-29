import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { JobList } from '@app/jobs/common/models/job.model';

import { CompaniesListModel, CompanyProfileModel } from '../models/company-profile.model';

@Injectable({
    providedIn: 'root'
})
export class CompanyProfileService {
    constructor(private readonly httpClient: HttpClient) {}

    getCompanyInfo(id: number): Observable<CompanyProfileModel> {
        return this.httpClient.get<CompanyProfileModel>(`${environment.API_BASE_URL}/api/v1/companies/${id}`);
    }

    getCompaniesList(): Observable<CompaniesListModel> {
        return this.httpClient.get<CompaniesListModel>(`${environment.API_BASE_URL}/api/v1/companies`);
    }

    deleteCompany(id: number): Observable<void> {
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/companies/${id}`);
    }

    getJobsByCompany(id: number): Observable<JobList> {
        return this.httpClient.get<JobList>(`${environment.API_BASE_URL}/api/v1/companies/${id}/jobs`);
    }
}
