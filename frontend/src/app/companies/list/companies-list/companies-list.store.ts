import { Observable, exhaustMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { CompaniesListModel } from '@app/companies/common/models/company-profile.model';
import { CompanyProfileService } from '@app/companies/common/services/company-profile.service';

export interface CompaniesListState {
    list: CompaniesListModel | null;
    loading: boolean;
    error: boolean;
}

const initialState: CompaniesListState = {
    list: null,
    loading: false,
    error: false
};

@Injectable()
export class CompaniesListStore extends ComponentStore<CompaniesListState> {
    vm$ = this.select({
        list: this.select((state) => state.list),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

    getList = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.getListLoading()),
            exhaustMap(() =>
                this.companiesService.getCompaniesList().pipe(
                    tapResponse(
                        (list) => this.getListLoadedSuccess(list),
                        () => this.getListLoadedError()
                    )
                )
            )
        )
    );

    private getListLoading = this.updater(
        (state): CompaniesListState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private getListLoadedSuccess = this.updater(
        (state, list: CompaniesListModel): CompaniesListState => ({
            ...state,
            list,
            loading: false
        })
    );
    private getListLoadedError = this.updater(
        (state): CompaniesListState => ({
            ...state,
            loading: false,
            error: true
        })
    );

    constructor(private readonly companiesService: CompanyProfileService) {
        super(initialState);
    }
}
