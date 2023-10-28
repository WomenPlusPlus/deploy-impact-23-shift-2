import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { JobLocationTypeEnum, JobTypeEnum } from '@app/common/models/jobs.model';
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
        // delete association from back end
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/companies/${id}`);
    }

    getJobsByCompany(id: number): Observable<JobList> {
        // get jobs with the specified Company ID
        //return this.httpClient.get<JobList>(`${environment.API_BASE_URL}/api/v1/jobs/company/${id}`);
        return of({
            items: [
                {
                    id: 1,
                    title: 'JUnior Software Developer',
                    jobType: JobTypeEnum.INTERNSHIP,
                    offerSalary: 85_000,
                    company: {
                        name: 'ZedTech',
                        imageUrl:
                            'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                        id: id
                    },
                    creator: {
                        imageUrl:
                            'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                        name: 'Test',
                        id: 2,
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
                    offerSalary: 75_000,
                    company: {
                        name: 'ZedTech',
                        imageUrl:
                            'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                        id: id
                    },
                    creator: {
                        imageUrl:
                            'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                        name: 'Test',
                        id: 2,
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
                    offerSalary: 80_000,
                    company: {
                        name: 'ZedTech',
                        imageUrl:
                            'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                        id: id
                    },
                    creator: {
                        imageUrl:
                            'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                        name: 'Test',
                        id: 2,
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
