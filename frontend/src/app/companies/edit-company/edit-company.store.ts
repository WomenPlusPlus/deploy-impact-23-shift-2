import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { CompanyItem } from '@app/companies/profile/common/models/company-profile.model';

import { CompaniesService } from '../common/services/companies.service';

export interface EditCompanyState {
    profile: CompanyItem | null;
    submitting: boolean;
    submitted: boolean;
}

const initialState: EditCompanyState = {
    profile: null,
    submitting: false,
    submitted: false
};

@Injectable()
export class EditCompanyStore extends ComponentStore<EditCompanyState> {
    profile$ = this.select((state) => state.profile);
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<FormData>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            exhaustMap((payload, id) =>
                this.companiesService.editCompany(payload, id).pipe(
                    tapResponse(
                        () => this.patchState({ submitting: false, submitted: true }),
                        () => {
                            this.toast.error(
                                "Could not update Company. Please try again later or contact the site's administrator."
                            );
                            this.patchState({ submitting: false, submitted: false });
                        }
                    )
                )
            )
        )
    );

    getValues(id: number): Observable<CompanyItem> {
        return this.companiesService.getCompany(id);
    }

    constructor(
        private readonly companiesService: CompaniesService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
