import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';

import { CreateUserFormModel } from './common/models/create-user.model';

export interface CreateUserState {
    response: { id: number } | null;
    submitting: boolean;
    submitted: boolean;
}

const initialState: CreateUserState = {
    response: null,
    submitting: false,
    submitted: false
};

@Injectable()
export class CreateUserStore extends ComponentStore<CreateUserState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted),
        response: this.select((state) => state.response)
    });

    submitForm = this.effect((trigger$: Observable<CreateUserFormModel>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false, response: null })),
            exhaustMap((payload) =>
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
