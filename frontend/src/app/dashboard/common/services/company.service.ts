import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { JobLocationTypeEnum, JobTypeEnum } from '@app/common/models/jobs.model';

import { JobListings } from '../models/jobs.model';
import { UsersList } from '../models/users.model';

@Injectable({
    providedIn: 'root'
})
export class CompanyDasboardService {
    constructor(private readonly httpClient: HttpClient) {}

    getJobsByCompany(id: number): Observable<JobListings> {
        //return this.httpClient.get<JobListings>(`${environment.API_BASE_URL}/api/v1/jobs/company/${id}`);
        return of({
            items: [
                {
                    id: 1,
                    title: 'JUnior Software Developer',
                    jobType: JobTypeEnum.INTERNSHIP,
                    creator: {
                        id: 2,
                        name: 'Test',
                        email: 'test@test.com'
                    },
                    company: {
                        id: id
                    },
                    creationDate: new Date(new Date().getTime() - 22 * 60 * 60 * 1000).toISOString(),
                    location: {
                        city: { id: 1, name: 'Zürich' },
                        type: JobLocationTypeEnum.HYBRID
                    }
                },
                {
                    id: 2,
                    title: 'UI/UX Designer',
                    jobType: JobTypeEnum.INTERNSHIP,
                    creator: {
                        id: 2,
                        name: 'Test',
                        email: 'test@test.com'
                    },
                    company: {
                        id: id
                    },
                    creationDate: new Date(new Date().getTime() - 22 * 60 * 60 * 1000).toISOString(),
                    location: {
                        city: { id: 1, name: 'Zürich' },
                        type: JobLocationTypeEnum.HYBRID
                    }
                },
                {
                    id: 3,
                    title: 'Project Manager',
                    jobType: JobTypeEnum.INTERNSHIP,
                    creator: {
                        id: 2,
                        name: 'Test',
                        email: 'test@test.com'
                    },
                    company: {
                        id: id
                    },
                    creationDate: new Date(new Date().getTime() - 22 * 60 * 60 * 1000).toISOString(),
                    location: {
                        city: { id: 1, name: 'Geneva' },
                        type: JobLocationTypeEnum.HYBRID
                    }
                }
            ]
        });
    }

    getUsersByCompany(): Observable<UsersList> {
        return this.httpClient.get<UsersList>(`${environment.API_BASE_URL}/api/v1/users`);
    }
}
