import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { AssociationProfileModel, AssociationsListModel } from '../models/association-profile.model';

@Injectable({
    providedIn: 'root'
})
export class AssociationProfileService {
    constructor(private readonly httpClient: HttpClient) {}

    getAssociationInfo(id: number): Observable<AssociationProfileModel> {
        //return this.httpClient.get<AssociationProfileModel>(`${environment.API_BASE_URL}/api/v1/associations/${id}`);
        return of({
            id: id,
            name: 'Test Association',
            imageUrl: {
                name: 'test',
                url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png'
            },
            websiteUrl: 'http://test-association-link',
            focus: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. '
        });
    }

    getAssociationsList(): Observable<AssociationsListModel> {
        return this.httpClient.get<AssociationsListModel>(`${environment.API_BASE_URL}/api/v1/admin/associations`);
        /*return of({
            items: [
                {
                    id: 0,
                    name: 'Test Association',
                    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    websiteUrl: 'http://test-association-link',
                    focus: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. '
                },
                {
                    id: 1,
                    name: 'Another Association',
                    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    websiteUrl: 'http://test-association-link',
                    focus: 'We are dedicated to innovation, collaboration, and excellence. We believe in pushing boundaries and creating cutting-edge solutions to drive positive change in the world.'
                },
                {
                    id: 2,
                    name: 'Monsters Inc.',
                    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    websiteUrl: 'http://test-association-link',
                    focus: 'Our core values revolve around technical excellence, customer-centricity, and a commitment to creating software that simplifies the complex.'
                },
                {
                    id: 3,
                    name: 'Dummy Association',
                    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    websiteUrl: 'http://test-association-link',
                    focus: 'At Dummy Association, we value innovation, sustainability, and diversity. We are dedicated to developing software that advances the world while protecting the environment.'
                }
            ]
        });*/
    }

    deleteAssociation(id: number): Observable<void> {
        // delete association from back end
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/admin/associations/${id}`);
    }
}
