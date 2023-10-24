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
        //return this.httpClient.get<CompanyProfileModel>(`${environment.API_BASE_URL}/api/v1/companies/${id}`);
        return of({
            id: id,
            name: 'Test Company',
            address: 'Street 123, Test',
            logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
            linkedin: 'testlink',
            kununu: 'testlink',
            phone: '0123456789',
            email: 'test@test.com',
            mission:
                'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
            values: 'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
            jobtypes:
                'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. ',
            expectation: 'Lorem ipsum dolor sit amet'
        });
    }

    getCompaniesList(): Observable<CompaniesListModel> {
        //return this.httpClient.get<CompanyProfileModel>(`${environment.API_BASE_URL}/api/v1/companies/`);
        return of({
            items: [
                {
                    id: 0,
                    name: 'Test Company',
                    address: 'Street 123, Test',
                    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    linkedin: 'https://www.linkedin.com/',
                    kununu: 'https://test-link-for-kununu',
                    phone: '0123456789',
                    email: 'test@test.com',
                    mission: 'Lorem ipsum dolor sit amet',
                    values: 'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
                    jobtypes:
                        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. ',
                    expectation: 'Lorem ipsum dolor sit amet'
                },
                {
                    id: 1,
                    name: 'Another Company',
                    address: 'Zurich, Switzerland',
                    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    linkedin: 'https://www.linkedin.com/',
                    kununu: 'https://test-link-for-kununu',
                    phone: '111222333',
                    email: 'test@test.com',
                    mission: 'Lorem ipsum dolor sit amet',
                    values: 'We are dedicated to innovation, collaboration, and excellence. We believe in pushing boundaries and creating cutting-edge solutions to drive positive change in the world.',
                    jobtypes:
                        'Join our team to work on exciting projects in AI, machine learning, and cloud computing.',
                    expectation:
                        "As part of our Company, you'll enjoy a flexible work environment, continuous learning opportunities, and a supportive team. We encourage creativity and provide a platform to make a real impact on the tech landscape."
                },
                {
                    id: 2,
                    name: 'Monsters Inc.',
                    address: 'Paris, France',
                    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    linkedin: 'https://www.linkedin.com/',
                    kununu: 'https://test-link-for-kununu',
                    phone: '0012300123',
                    email: 'test@test.com',
                    mission: 'Lorem ipsum dolor sit amet',
                    values: 'Our core values revolve around technical excellence, customer-centricity, and a commitment to creating software that simplifies the complex.',
                    jobtypes: 'Join us to work on exciting projects in web and mobile app development.',
                    expectation:
                        "As a part of this company, you can expect a dynamic and inspiring work environment. We foster a culture of growth and collaboration, and we're dedicated to helping our employees reach their full potential."
                },
                {
                    id: 3,
                    name: 'Dummy Company',
                    address: 'Berlin, Germany',
                    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    linkedin: 'https://www.linkedin.com/',
                    kununu: 'https://test-link-for-kununu',
                    phone: '1234512345',
                    email: 'test@test.com',
                    mission: 'Lorem ipsum dolor sit amet',
                    values: 'At Dummy Company, we value innovation, sustainability, and diversity. We are dedicated to developing software that advances the world while protecting the environment.',
                    jobtypes: 'Join us to work on groundbreaking projects in green tech. ',
                    expectation:
                        'Dummy Company offers a unique opportunity to make a positive impact on the environment. We provide a stimulating work atmosphere, ongoing support for your career development, and the chance to be part of a company with a global mission for a sustainable future.'
                }
            ]
        });
    }

    deleteCompany(id: number): Observable<void> {
        // delete association from back end
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/admin/companies/${id}`);
    }

    getJobsByCompany(id: number): Observable<JobList> {
        // get jobs with the specified Company ID
        //return this.httpClient.get<JobList>(`${environment.API_BASE_URL}/api/v1/jobs/${id}`);
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
