import { HotToastService } from '@ngneat/hot-toast';
import { catchError, EMPTY, Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminCompanyService } from '@app/admin/company/common/services/admin-company.service';
import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';
import { UserFormCompanyFormModel, UserFormSubmissionModel } from '@app/admin/users/form/common/models/user-form.model';

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
                this.adminCompanyService.createCompany(payload.company).pipe(
                    switchMap(({ id }) =>
                        this.adminUsersService
                            .setupUser({ ...payload.user, companyId: id } as UserFormCompanyFormModel)
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
}