import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Profile } from '@app/common/models/profile.model';
import environment from '@envs/environment';

@Injectable({
    providedIn: 'root'
})
export class AccountService {

    constructor(private readonly httpClient: HttpClient) {
    }

    me(): Observable<Profile> {
        return this.httpClient.get<Profile>(`${environment.API_BASE_URL}/api/v1/me`);
    }

}
