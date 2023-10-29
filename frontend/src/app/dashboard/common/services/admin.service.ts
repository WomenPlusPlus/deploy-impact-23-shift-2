import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { JobList } from '@app/jobs/common/models/job.model';
import { UsersList } from '@app/users/common/models/users-list.model';

import { AssociationList } from '../models/associations.model';
import { CompanyList } from '../models/company.model';

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

    getJobs(): Observable<JobList> {
        return this.httpClient.get<JobList>(`${environment.API_BASE_URL}/api/v1/jobs`);
    }
}
