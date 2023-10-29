import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { JobList } from '@app/jobs/common/models/job.model';

import { UsersList } from '../models/users.model';

@Injectable({
    providedIn: 'root'
})
export class CompanyDasboardService {
    constructor(private readonly httpClient: HttpClient) {}

    getJobsByCompany(id: number): Observable<JobList> {
        return this.httpClient.get<JobList>(`${environment.API_BASE_URL}/api/v1/companies/${id}/jobs`);
    }

    getUsersByCompany(): Observable<UsersList> {
        return this.httpClient.get<UsersList>(`${environment.API_BASE_URL}/api/v1/users`);
    }
}
