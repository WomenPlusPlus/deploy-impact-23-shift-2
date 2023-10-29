import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { AssociationList } from '../models/associations.model';
import { CompanyList } from '../models/company.model';
import { JobListings } from '../models/jobs.model';
import { UsersList } from '../models/users.model';

@Injectable({
    providedIn: 'root'
})
export class AdminDashboardService {
    constructor(private readonly httpClient: HttpClient) {}

    getUsers(): Observable<UsersList> {
        return this.httpClient.get<UsersList>(`${environment.API_BASE_URL}/api/v1/users`);
    }

    getCompanies(): Observable<CompanyList> {
        return this.httpClient.get<CompanyList>(`${environment.API_BASE_URL}/api/v1/companies`);
    }

    getAssociations(): Observable<AssociationList> {
        return this.httpClient.get<AssociationList>(`${environment.API_BASE_URL}/api/v1/associations`);
    }

    getJobs(): Observable<JobListings> {
        return this.httpClient.get<JobListings>(`${environment.API_BASE_URL}/api/v1/jobs`);
    }
}
