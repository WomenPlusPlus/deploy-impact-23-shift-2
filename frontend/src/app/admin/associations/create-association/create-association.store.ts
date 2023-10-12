import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AdminAssociationService } from '../common/services/admin-association.service';

export interface CreateAssociationState {
    submitting: boolean;
    submitted: boolean;
}

const initialState: CreateAssociationState = {
    submitting: false,
    submitted: false
};

@Injectable()
export class CreateAssociationStore extends ComponentStore<CreateAssociationState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<FormData>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            exhaustMap((payload) =>
                this.adminAssociationService.createAssociation(payload).pipe(
                    tapResponse(
                        () => this.patchState({ submitting: false, submitted: true }),
                        () => {
                            this.toast.error(
                                "Could not create Association. Please try again later or contact the site's administrator."
                            );
                            this.patchState({ submitting: false, submitted: false });
                        }
                    )
                )
            )
        )
    );

    constructor(
        private readonly adminAssociationService: AdminAssociationService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
