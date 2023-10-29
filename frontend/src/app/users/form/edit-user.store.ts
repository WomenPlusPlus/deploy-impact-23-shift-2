import { HotToastService } from '@ngneat/hot-toast';
import { Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { UserDetails } from '@app/common/models/users.model';
import { AdminUsersService } from '@app/users/common/services/admin-users.service';
import { EditUserResponse } from '@app/users/form/common/models/edit-user.model';

import { UserFormSubmissionModel } from './common/models/user-form.model';

export interface UserFormState {
    user: UserDetails | null;
    response: EditUserResponse | null;
    submitting: boolean;
    submitted: boolean;
}

const initialState: UserFormState = {
    user: null,
    response: null,
    submitting: false,
    submitted: false
};

@Injectable()
export class EditUserStore extends ComponentStore<UserFormState> {
    user$ = this.select((state) => state.user);
    vm$ = this.select({
        user: this.user$,
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted),
        response: this.select((state) => state.response)
    });

    submitForm = this.effect((trigger$: Observable<UserFormSubmissionModel & { id: number }>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false, response: null })),
            switchMap((payload) =>
                this.adminUsersService.editUser(payload.id, payload).pipe(
                    tapResponse(
                        (response) => this.patchState({ submitting: false, submitted: true, response }),
                        () => {
                            this.toast.error(
                                'Could not edit a new user! Please try again later or contact the support.'
                            );
                            this.patchState({ submitting: false, submitted: false });
                        }
                    )
                )
            )
        )
    );

    getUser = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            switchMap((id: number) =>
                this.adminUsersService.getById(id).pipe(
                    tapResponse(
                        (user) => this.patchState({ user }),
                        () => {
                            this.toast.error('Could not load the user! Please try again later or contact the support.');
                            this.router.navigate(['..']);
                        }
                    )
                )
            )
        )
    );

    constructor(
        private readonly router: Router,
        private readonly adminUsersService: AdminUsersService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
