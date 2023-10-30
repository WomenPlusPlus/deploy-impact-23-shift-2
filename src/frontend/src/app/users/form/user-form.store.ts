import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AssociationsService } from '@app/associations/common/services/associations.service';
import { CompaniesService } from '@app/companies/common/services/companies.service';
import { CompanyItem } from '@app/companies/profile/common/models/company-profile.model';
import { AssociationItem } from '@app/dashboard/common/models/associations.model';

export interface UserFormState {
    loadingCompanies: boolean;
    loadedCompanies: boolean;
    errorCompanies: boolean;
    companies: CompanyItem[];
    loadingAssociations: boolean;
    loadedAssociations: boolean;
    errorAssociations: boolean;
    associations: AssociationItem[];
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
                this.companiesService.getCompaniesList().pipe(
                    tapResponse(
                        (list) =>
                            this.patchState({
                                loadingCompanies: false,
                                loadedCompanies: true,
                                companies: list.items
                            }),
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
                this.associationsService.getAssociationsList().pipe(
                    tapResponse(
                        (list) =>
                            this.patchState({
                                loadingAssociations: false,
                                loadedAssociations: true,
                                associations: list.items
                            }),
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
        private readonly companiesService: CompaniesService,
        private readonly associationsService: AssociationsService
    ) {
        super(initialState);
    }
}
