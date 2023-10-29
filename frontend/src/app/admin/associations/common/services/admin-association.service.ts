import { Observable } from 'rxjs';

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
        return this.httpClient.post<{ id: number }>(`${environment.API_BASE_URL}/api/v1/associations`, payload);
    }

    editAssociation(payload: FormData, id: number): Observable<void> {
        return this.httpClient.put<void>(`${environment.API_BASE_URL}/api/v1/associations/${id}`, payload);
    }

    getAssociation(id: number): Observable<AssociationProfileModel> {
        return this.httpClient.get<AssociationProfileModel>(`${environment.API_BASE_URL}/api/v1/associations/${id}`);
    }
}
