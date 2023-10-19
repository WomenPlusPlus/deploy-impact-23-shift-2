import { HotToastService } from '@ngneat/hot-toast';
import { Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';

import { UserFormSubmissionModel } from './common/models/user-form.model';
import { CreateUserResponse } from '@app/admin/users/form/common/models/create-user.model';

export interface UserFormState {
    response: CreateUserResponse | null;
    submitting: boolean;
    submitted: boolean;
}

const initialState: UserFormState = {
    response: null,
    submitting: false,
    submitted: false
};

@Injectable()
export class CreateUserStore extends ComponentStore<UserFormState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted),
        response: this.select((state) => state.response)
    });

    submitForm = this.effect((trigger$: Observable<UserFormSubmissionModel>) =>
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

    constructor(
        private readonly adminUsersService: AdminUsersService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
