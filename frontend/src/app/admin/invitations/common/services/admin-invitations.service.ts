import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { CreateInviteFormModel } from '@app/admin/invitations/create-invite/common/models/create-invite.model';

@Injectable({
    providedIn: 'root'
})
export class AdminInvitationsService {
    constructor(private readonly httpClient: HttpClient) {}

    invite(payload: CreateInviteFormModel): Observable<void> {
        return this.httpClient.post<void>(`${environment.API_BASE_URL}/api/v1/admin/invitations`, payload);
    }
}