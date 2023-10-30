import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AssociationsService } from '@app/associations/common/services/associations.service';
import { CompaniesService } from '@app/companies/common/services/companies.service';
import { CompanyItem } from '@app/companies/profile/common/models/company-profile.model';
import { AssociationItem } from '@app/dashboard/common/models/associations.model';
import { InvitationsService } from '@app/invitations/common/services/invitations.service';
import { CreateInviteFormModel } from '@app/invitations/create-invite/common/models/create-invite.model';

export interface CreateInviteState {
    submitting: boolean;
    submitted: boolean;
    companies: CompanyItem[];
    associations: AssociationItem[];
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
                this.invitationsService.invite(payload).pipe(
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
                this.companiesService.getCompaniesList().pipe(
                    tapResponse(
                        (list) => this.patchState({ companies: list.items }),
                        (error) => console.error('Could not load companies: ', error)
                    )
                )
            )
        )
    );

    loadAssociations = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            exhaustMap(() =>
                this.associationsService.getAssociationsList().pipe(
                    tapResponse(
                        (list) => this.patchState({ associations: list.items }),
                        (error) => console.error('Could not load associations: ', error)
                    )
                )
            )
        )
    );

    constructor(
        private readonly invitationsService: InvitationsService,
        private readonly companiesService: CompaniesService,
        private readonly associationsService: AssociationsService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
