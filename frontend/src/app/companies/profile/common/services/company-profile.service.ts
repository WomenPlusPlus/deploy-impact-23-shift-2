import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { CompaniesListModel, CompanyProfileModel } from '../models/company-profile.model';

//import environment from '@envs/environment';

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
                    values: 'At Dummy Company, we value innovation, sustainability, and diversity. We are dedicated to developing software that advances the world while protecting the environment.',
                    jobtypes: 'Join us to work on groundbreaking projects in green tech. ',
                    expectation:
                        'Dummy Company offers a unique opportunity to make a positive impact on the environment. We provide a stimulating work atmosphere, ongoing support for your career development, and the chance to be part of a company with a global mission for a sustainable future.'
                }
            ]
        });
    }
}
