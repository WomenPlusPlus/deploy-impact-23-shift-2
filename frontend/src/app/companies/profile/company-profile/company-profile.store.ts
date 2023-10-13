import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { CompanyProfileModel } from '../common/models/company-profile.model';
import { CompanyProfileService } from '../common/services/company-profile.service';

export interface CompanyProfileState {
    profile: CompanyProfileModel | null;
    loading: boolean;
    error: boolean;
}

const initialState: CompanyProfileState = {
    profile: null,
    loading: false,
    error: false
};

@Injectable()
export class CompanyProfileStore extends ComponentStore<CompanyProfileState> {
    vm$ = this.select({
        profile: this.select((state) => state.profile),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

    getProfile = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.profileLoading()),
            exhaustMap((id: number) =>
                this.companyProfileService.getCompanyInfo(id).pipe(
                    tapResponse(
                        (profile) => this.profileLoadSuccess(profile),
                        () => this.profileLoadError()
                    )
                )
            )
        )
    );

    private profileLoading = this.updater(
        (state): CompanyProfileState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private profileLoadSuccess = this.updater(
        (state, profile: CompanyProfileModel): CompanyProfileState => ({
            ...state,
            profile,
            loading: false
        })
    );
    private profileLoadError = this.updater(
        (state): CompanyProfileState => ({
            ...state,
            error: true,
            loading: false
        })
    );

    constructor(private readonly companyProfileService: CompanyProfileService) {
        super(initialState);
    }
}
