import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { UsersList } from '../models/users.model';

@Injectable({
    providedIn: 'root'
})
export class AssociationDasboardService {
    constructor(private readonly httpClient: HttpClient) {}

    getUsersByAssociation(): Observable<UsersList> {
        // get users with the specified Association ID
        return this.httpClient.get<UsersList>(`${environment.API_BASE_URL}/api/v1/users`);
    }
}
