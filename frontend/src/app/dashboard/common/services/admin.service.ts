import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { JobTypeEnum, JobLocationTypeEnum } from '@app/common/models/jobs.model';

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
        //return this.httpClient.get<CompanyList>(`${environment.API_BASE_URL}/api/v1/companies/`);
        return of({
            items: [
                {
                    id: 0,
                    name: 'Test Company',
                    email: 'test@test.com'
                },
                {
                    id: 1,
                    name: 'Another Company',
                    email: 'test@test.com'
                },
                {
                    id: 2,
                    name: 'Monsters Inc.',
                    email: 'test@test.com'
                },
                {
                    id: 3,
                    name: 'Dummy Company',
                    email: 'test@test.com'
                }
            ]
        });
    }

    getAssociations(): Observable<AssociationList> {
        return this.httpClient.get<AssociationList>(`${environment.API_BASE_URL}/api/v1/admin/associations`);
    }

    getJobs(): Observable<JobListings> {
        // get jobs with the specified Company ID
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
                    creationDate: new Date(new Date().getTime() - 22 * 60 * 60 * 1000).toISOString(),
                    location: {
                        city: { id: 1, name: 'Geneva' },
                        type: JobLocationTypeEnum.HYBRID
                    }
                }
            ]
        });
    }
}
