import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { AssociationProfileModel } from '@app/associations/common/models/association-profile.model';
import { AssociationsService } from '@app/associations/common/services/associations.service';

export interface AssociationProfileState {
    profile: AssociationProfileModel | null;
    loading: boolean;
    error: boolean;
}

const initialState: AssociationProfileState = {
    profile: null,
    loading: false,
    error: false
};

@Injectable()
export class AssociationProfileStore extends ComponentStore<AssociationProfileState> {
    vm$ = this.select({
        profile: this.select((state) => state.profile),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

    getProfile = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.profileLoading()),
            exhaustMap((id: number) =>
                this.associationsService.getAssociation(id).pipe(
                    tapResponse(
                        (profile) => this.profileLoadSuccess(profile),
                        () => {
                            this.profileLoadError();
                            this.router.navigate(['/associations']);
                            this.toast.error('Association with id ' + id + ' not found.');
                        }
                    )
                )
            )
        )
    );

    private profileLoading = this.updater(
        (state): AssociationProfileState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private profileLoadSuccess = this.updater(
        (state, profile: AssociationProfileModel): AssociationProfileState => ({
            ...state,
            profile,
            loading: false
        })
    );
    private profileLoadError = this.updater(
        (state): AssociationProfileState => ({
            ...state,
            error: true,
            loading: false
        })
    );

    constructor(
        private readonly associationsService: AssociationsService,
        private router: Router,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
