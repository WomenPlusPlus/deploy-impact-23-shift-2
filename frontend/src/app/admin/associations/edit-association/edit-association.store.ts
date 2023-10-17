import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AssociationProfileModel } from '@app/associations/common/models/association-profile.model';

import { AdminAssociationService } from '../common/services/admin-association.service';

export interface EditAssociationState {
    profile: AssociationProfileModel | null;
    submitting: boolean;
    submitted: boolean;
}

const initialState: EditAssociationState = {
    profile: null,
    submitting: false,
    submitted: false
};

@Injectable()
export class EditAssociationStore extends ComponentStore<EditAssociationState> {
    profile$ = this.select((state) => state.profile);
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<FormData>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            exhaustMap((payload, id) =>
                this.adminAssociationService.editAssociation(payload, id).pipe(
                    tapResponse(
                        () => this.patchState({ submitting: false, submitted: true }),
                        () => {
                            this.toast.error(
                                "Could not update Association. Please try again later or contact the site's administrator."
                            );
                            this.patchState({ submitting: false, submitted: false });
                        }
                    )
                )
            )
        )
    );

    getValues(id: number): Observable<AssociationProfileModel> {
        return this.adminAssociationService.getAssociation(id);
    }

    constructor(
        private readonly adminAssociationService: AdminAssociationService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
