import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminInvitationsService } from '@app/admin/invitations/common/services/admin-invitations.service';
import { CreateInviteFormModel } from '@app/admin/invitations/create-invite/common/models/create-invite.model';
import { Association } from '@app/common/models/associations.model';
import { Company } from '@app/common/models/companies.model';
import { AdminAssociationsService } from '@app/users/common/services/admin-associations.service';
import { AdminCompaniesService } from '@app/users/common/services/admin-companies.service';

export interface CreateInviteState {
    submitting: boolean;
    submitted: boolean;
    companies: Company[];
    associations: Association[];
}

const initialState: CreateInviteState = {
    submitting: false,
    submitted: false,
    companies: [],
    associations: []
};

@Injectable()
export class CreateInviteStore extends ComponentStore<CreateInviteState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted),
        companies: this.select((state) => state.companies),
        associations: this.select((state) => state.associations)
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

    loadData = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(switchMap(() => [this.loadCompanies(), this.loadAssociations()]))
    );

    loadCompanies = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            exhaustMap(() =>
                this.adminCompaniesService.getCompanies().pipe(
                    tapResponse(
                        (companies) => this.patchState({ companies }),
                        (error) => console.error('Could not load companies: ', error)
                    )
                )
            )
        )
    );

    loadAssociations = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            exhaustMap(() =>
                this.adminAssociationsService.getAssociations().pipe(
                    tapResponse(
                        (associations) => this.patchState({ associations }),
                        (error) => console.error('Could not load associations: ', error)
                    )
                )
            )
        )
    );

    constructor(
        private readonly adminInvitationsService: AdminInvitationsService,
        private readonly adminCompaniesService: AdminCompaniesService,
        private readonly adminAssociationsService: AdminAssociationsService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
