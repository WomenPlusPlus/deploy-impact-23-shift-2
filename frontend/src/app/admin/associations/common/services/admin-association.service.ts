import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

@Injectable({
    providedIn: 'root'
})
export class AdminAssociationService {
    constructor(private readonly httpClient: HttpClient) {}

    createAssociation(payload: FormData): Observable<void> {
        // post association data to back end
        return this.httpClient.post<void>(`${environment.API_BASE_URL}/api/v1/admin/associations`, payload);
    }
}
