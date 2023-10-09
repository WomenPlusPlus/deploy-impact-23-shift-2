import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminInvitationsService } from '@app/admin/invitations/common/services/admin-invitations.service';
import { CreateInviteFormModel } from '@app/admin/invitations/create-invite/common/models/create-invite.model';

export interface CreateInviteState {
    submitting: boolean;
    submitted: boolean;
}

const initialState: CreateInviteState = {
    submitting: false,
    submitted: false
};

@Injectable()
export class CreateInviteStore extends ComponentStore<CreateInviteState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<CreateInviteFormModel>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            exhaustMap((payload) =>
                this.adminInvitationsService.invite(payload).pipe(
                    tapResponse(
                        () => this.patchState({ submitting: false, submitted: true }),
                        () => {
                            this.toast.error('Could not invite user! Please try again later or contact the support.');
                            this.patchState({ submitting: false, submitted: false });
                        }
                    )
                )
            )
        )
    );

    constructor(
        private readonly adminInvitationsService: AdminInvitationsService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
