import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminAssociationsService } from '@app/admin/users/common/services/admin-associations.service';
import { AdminCompaniesService } from '@app/admin/users/common/services/admin-companies.service';
import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';
import { Association } from '@app/common/models/associations.model';
import { Company } from '@app/common/models/companies.model';

import { CreateUserResponse, CreateUserSubmissionModel } from './common/models/create-user.model';

export interface CreateUserState {
    response: CreateUserResponse | null;
    submitting: boolean;
    submitted: boolean;
    loadingCompanies: boolean;
    loadedCompanies: boolean;
    errorCompanies: boolean;
    companies: Company[];
    loadingAssociations: boolean;
    loadedAssociations: boolean;
    errorAssociations: boolean;
    associations: Association[];
}

const initialState: CreateUserState = {
    response: null,
    submitting: false,
    submitted: false,
    loadingCompanies: false,
    loadedCompanies: false,
    errorCompanies: false,
    companies: [],
    loadingAssociations: false,
    loadedAssociations: false,
    errorAssociations: false,
    associations: []
};

@Injectable()
export class CreateUserStore extends ComponentStore<CreateUserState> {
    companies$ = this.select((state) => state.companies);
    associations$ = this.select((state) => state.associations);
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted),
        response: this.select((state) => state.response)
    });

    submitForm = this.effect((trigger$: Observable<CreateUserSubmissionModel>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false, response: null })),
            switchMap((payload) =>
                this.adminUsersService.createUser(payload).pipe(
                    tapResponse(
                        (response) => this.patchState({ submitting: false, submitted: true, response }),
                        () => {
                            this.toast.error(
                                'Could not create a new user! Please try again later or contact the support.'
                            );
                            this.patchState({ submitting: false, submitted: false });
                        }
                    )
                )
            )
        )
    );

    loadCompanies = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.patchState({ loadingCompanies: true, loadedCompanies: false, errorCompanies: false })),
            exhaustMap(() =>
                this.adminCompaniesService.getCompanies().pipe(
                    tapResponse(
                        (companies) => this.patchState({ loadingCompanies: false, loadedCompanies: true, companies }),
                        (error) => {
                            console.error('Could not load companies: ', error);
                            this.patchState({ loadingCompanies: false, errorCompanies: true });
                        }
                    )
                )
            )
        )
    );

    loadAssociations = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() =>
                this.patchState({ loadingAssociations: true, loadedAssociations: false, errorAssociations: false })
            ),
            exhaustMap(() =>
                this.adminAssociationsService.getAssociations().pipe(
                    tapResponse(
                        (associations) =>
                            this.patchState({ loadingAssociations: false, loadedAssociations: true, associations }),
                        (error) => {
                            console.error('Could not load associations: ', error);
                            this.patchState({ loadingAssociations: false, errorAssociations: true });
                        }
                    )
                )
            )
        )
    );

    constructor(
        private readonly adminUsersService: AdminUsersService,
        private readonly adminCompaniesService: AdminCompaniesService,
        private readonly adminAssociationsService: AdminAssociationsService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
