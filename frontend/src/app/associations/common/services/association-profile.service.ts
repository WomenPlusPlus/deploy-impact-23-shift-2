import { Observable } from 'rxjs';

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
        return this.httpClient.get<AssociationProfileModel>(`${environment.API_BASE_URL}/api/v1/associations/${id}`);
    }

    getAssociationsList(): Observable<AssociationsListModel> {
        return this.httpClient.get<AssociationsListModel>(`${environment.API_BASE_URL}/api/v1/associations`);
    }

    deleteAssociation(id: number): Observable<void> {
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/associations/${id}`);
    }
}
