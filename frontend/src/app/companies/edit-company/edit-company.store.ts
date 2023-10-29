import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { CompanyProfileModel } from '@app/companies/profile/common/models/company-profile.model';

import { AdminCompanyService } from '../common/services/admin-company.service';

export interface EditCompanyState {
    profile: CompanyProfileModel | null;
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
                this.adminCompanyService.editCompany(payload, id).pipe(
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

    getValues(id: number): Observable<CompanyProfileModel> {
        return this.adminCompanyService.getCompany(id);
    }

    constructor(
        private readonly adminCompanyService: AdminCompanyService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
