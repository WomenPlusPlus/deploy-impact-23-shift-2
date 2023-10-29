import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { InvitationList } from '@app/common/models/invitation.model';
import { CreateInviteFormModel } from '@app/invitations/create-invite/common/models/create-invite.model';

@Injectable({
    providedIn: 'root'
})
export class AdminInvitationsService {
    constructor(private readonly httpClient: HttpClient) {}

    invite(payload: CreateInviteFormModel): Observable<void> {
        return this.httpClient.post<void>(`${environment.API_BASE_URL}/api/v1/invitations`, payload);
    }

    getList(): Observable<InvitationList> {
        return this.httpClient.get<InvitationList>(`${environment.API_BASE_URL}/api/v1/invitations`);
    }
}
