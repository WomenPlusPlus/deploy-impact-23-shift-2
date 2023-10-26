import { HotToastService } from '@ngneat/hot-toast';
import { Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';
import { UserFormSubmissionModel } from '@app/admin/users/form/common/models/user-form.model';

export interface UserFormState {
    submitting: boolean;
    submitted: boolean;
}

const initialState: UserFormState = {
    submitting: false,
    submitted: false
};

@Injectable()
export class SetupCandidateFormStore extends ComponentStore<UserFormState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<UserFormSubmissionModel>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            switchMap((payload) =>
                this.adminUsersService.setupUser(payload).pipe(
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
            )
        )
    );

    constructor(
        private readonly adminUsersService: AdminUsersService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
