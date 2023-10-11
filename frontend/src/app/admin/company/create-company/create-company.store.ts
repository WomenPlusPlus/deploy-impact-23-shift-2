import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminCompanyService } from '../common/services/admin-company.service';
import { CreateCompanyFormModel } from './common/models/create-company.model';

export interface CreateCompanyState {
    submitting: boolean;
    submitted: boolean;
}

const initialState: CreateCompanyState = {
    submitting: false,
    submitted: false
};

@Injectable()
export class CreateCompanyStore extends ComponentStore<CreateCompanyState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<CreateCompanyFormModel>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            exhaustMap((payload) =>
                this.adminCompanyService.createCompany(payload).pipe(
                    tapResponse(
                        () => this.patchState({ submitting: false, submitted: true }),
                        () => {
                            this.toast.error(
                                "Could not create company. Please try again later or contact the site's administrator."
                            );
                            this.patchState({ submitting: false, submitted: false });
                        }
                    )
                )
            )
        )
    );

    constructor(
        private readonly adminCompanyService: AdminCompanyService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
