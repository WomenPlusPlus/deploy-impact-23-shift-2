import { HotToastService } from '@ngneat/hot-toast';
import { Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { ProfileSetup } from '@app/common/models/profile.model';
import { AdminUsersService } from '@app/users/common/services/admin-users.service';

export interface SetupFormState {
    profile: ProfileSetup | null;
    loading: boolean;
    loaded: boolean;
    error: boolean;
}

const initialState: SetupFormState = {
    profile: null,
    loading: false,
    loaded: false,
    error: false
};

@Injectable()
export class SetupScreenStore extends ComponentStore<SetupFormState> {
    vm$ = this.select({
        profile: this.select((state) => state.profile),
        loading: this.select((state) => state.loading),
        loaded: this.select((state) => state.loaded),
        error: this.select((state) => state.error)
    });

    getProfile = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.patchState({ loading: true, loaded: false, error: false })),
            switchMap(() =>
                this.adminUsersService.getSetupInfo().pipe(
                    tapResponse(
                        (profile) => this.patchState({ profile, loading: false, loaded: true }),
                        () => {
                            this.toast.error(
                                'Could not load the user setup info! Please try again later or contact the support.'
                            );
                            this.patchState({ loading: false, loaded: true, error: true });
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
