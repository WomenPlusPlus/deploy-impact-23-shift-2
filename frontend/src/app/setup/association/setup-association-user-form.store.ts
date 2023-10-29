import { HotToastService } from '@ngneat/hot-toast';
import { catchError, EMPTY, map, Observable, of, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AssociationsService } from '@app/associations/common/services/associations.service';
import { UsersService } from '@app/users/common/services/users.service';
import { UserFormAssociationFormModel, UserFormSubmissionModel } from '@app/users/form/common/models/user-form.model';

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
    associationId?: number;
    association: FormData;
}

@Injectable()
export class SetupAssociationUserFormStore extends ComponentStore<UserFormState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<FormSubmissionModel>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            switchMap((payload) =>
                this.createAssociation(payload).pipe(
                    switchMap((associationId) =>
                        this.adminUsersService
                            .setupUser({ ...payload.user, associationId } as UserFormAssociationFormModel)
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
                            'Could not setup your association! Please try again later or contact the support.'
                        );
                        this.patchState({ submitting: false, submitted: false });
                        return EMPTY;
                    })
                )
            )
        )
    );

    constructor(
        private readonly adminUsersService: UsersService,
        private readonly associationsService: AssociationsService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }

    private createAssociation({ associationId, association }: FormSubmissionModel): Observable<number> {
        if (associationId) {
            return of(associationId);
        }
        return this.associationsService.createAssociation(association).pipe(map(({ id }) => id));
    }
}
