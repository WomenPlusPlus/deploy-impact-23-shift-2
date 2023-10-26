import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { AssociationProfileModel } from '@app/associations/common/models/association-profile.model';

@Injectable({
    providedIn: 'root'
})
export class AdminAssociationService {
    constructor(private readonly httpClient: HttpClient) {}

    createAssociation(payload: FormData): Observable<{ id: number }> {
        // post association data to back end
        return this.httpClient.post<{ id: number }>(`${environment.API_BASE_URL}/api/v1/admin/associations`, payload);
    }

    editAssociation(payload: FormData, id: number): Observable<void> {
        // update association data to back end
        return this.httpClient.put<void>(`${environment.API_BASE_URL}/api/v1/admin/associations/${id}`, payload);
    }

    getAssociation(id: number): Observable<AssociationProfileModel> {
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

    deleteAssociation(id: number): Observable<void> {
        // post association data to back end
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/admin/associations/${id}`);
    }
}
