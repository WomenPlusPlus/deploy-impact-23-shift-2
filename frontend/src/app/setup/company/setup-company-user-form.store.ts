import { HotToastService } from '@ngneat/hot-toast';
import { catchError, EMPTY, map, Observable, of, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminCompanyService } from '@app/companies/common/services/admin-company.service';
import { AdminUsersService } from '@app/users/common/services/admin-users.service';
import { UserFormCompanyFormModel, UserFormSubmissionModel } from '@app/users/form/common/models/user-form.model';

export interface UserFormState {
    submitting: boolean;
    submitted: boolean;
}

const initialState: UserFormState = {
    submitting: false,
    submitted: false
};

interface FormSubmissionModel {
    user: UserFormSubmissionModel;
    companyId?: number;
    company: FormData;
}

@Injectable()
export class SetupCompanyUserFormStore extends ComponentStore<UserFormState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<FormSubmissionModel>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            switchMap((payload) =>
                this.createCompany(payload).pipe(
                    switchMap((companyId) =>
                        this.adminUsersService
                            .setupUser({ ...payload.user, companyId } as UserFormCompanyFormModel)
                            .pipe(
                                tapResponse(
                                    () => {
                                        this.patchState({ submitting: false, submitted: true });
                                        window.location.reload();
                                    },
                                    () => {
                                        this.toast.error(
                                            'Could not setup your profile! Please try again later or contact the support.'
                                        );
                                        this.patchState({ submitting: false, submitted: false });
                                    }
                                )
                            )
                    ),
                    catchError(() => {
                        this.toast.error(
                            'Could not setup your company! Please try again later or contact the support.'
                        );
                        this.patchState({ submitting: false, submitted: false });
                        return EMPTY;
                    })
                )
            )
        )
    );

    constructor(
        private readonly adminUsersService: AdminUsersService,
        private readonly adminCompanyService: AdminCompanyService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }

    private createCompany({ companyId, company }: FormSubmissionModel): Observable<number> {
        if (companyId) {
            return of(companyId);
        }
        return this.adminCompanyService.createCompany(company).pipe(map(({ id }) => id));
    }
}
