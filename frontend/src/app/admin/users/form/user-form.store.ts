import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminAssociationsService } from '@app/admin/users/common/services/admin-associations.service';
import { AdminCompaniesService } from '@app/admin/users/common/services/admin-companies.service';
import { Association } from '@app/common/models/associations.model';
import { Company } from '@app/common/models/companies.model';

export interface UserFormState {
    loadingCompanies: boolean;
    loadedCompanies: boolean;
    errorCompanies: boolean;
    companies: Company[];
    loadingAssociations: boolean;
    loadedAssociations: boolean;
    errorAssociations: boolean;
    associations: Association[];
}

const initialState: UserFormState = {
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
export class UserFormStore extends ComponentStore<UserFormState> {
    companies$ = this.select((state) => state.companies);
    associations$ = this.select((state) => state.associations);

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
        private readonly adminCompaniesService: AdminCompaniesService,
        private readonly adminAssociationsService: AdminAssociationsService
    ) {
        super(initialState);
    }
}
