import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

@Injectable({
    providedIn: 'root'
})
export class AdminCompanyService {
    constructor(private readonly httpClient: HttpClient) {}

    createCompany(payload: FormData): Observable<void> {
        // post company data to back
        return this.httpClient.post<void>(`${environment.API_BASE_URL}/api/v1/admin/company`, payload);
    }
}
