import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { CompanyProfileModel } from '../models/company-profile.model';

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
            mission:
                'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
            values: 'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
            jobtypes:
                'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. ',
            expectation: 'Lorem ipsum dolor sit amet'
        });
    }
}
