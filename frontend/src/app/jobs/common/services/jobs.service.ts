import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { Job, JobList } from '@app/jobs/common/models/job.model';

//import environment from '@envs/environment';

@Injectable({
    providedIn: 'root'
})
export class JobsService {
    constructor(private readonly httpClient: HttpClient) {}

    getJobDetails(id: number): Observable<Job> {
        return this.httpClient.get<Job>(`${environment.API_BASE_URL}/api/v1/jobs/${id}`);
    }

    getList(): Observable<JobList> {
        return this.httpClient.get<JobList>(`${environment.API_BASE_URL}/api/v1/jobs`);
    }
}
