import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { Profile } from '@app/common/models/profile.model';

@Injectable({
    providedIn: 'root'
})
export class AccountService {
    constructor(private readonly httpClient: HttpClient) {}

    me(): Observable<Profile> {
        return this.httpClient.get<Profile>(`${environment.API_BASE_URL}/api/v1/me`);
    }
}
