import Fuse from 'fuse.js';
import { pick } from 'lodash';
import { Observable, exhaustMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { CompaniesListModel } from '@app/companies/common/models/company-profile.model';
import { CompanyProfileService } from '@app/companies/common/services/company-profile.service';

export interface CompaniesListState {
    list: CompaniesListModel | null;
    loading: boolean;
    error: boolean;
    searchString: string;
}

const initialState: CompaniesListState = {
    list: null,
    loading: false,
    error: false,
    searchString: ''
};

const SEARCH_TERM_MIN_LEN = 3;

@Injectable()
export class CompaniesListStore extends ComponentStore<CompaniesListState> {
    private list$ = this.select((state) => state.list);
    searchString$ = this.select((state) => state.searchString);

    vm$ = this.select({
        list: this.select(this.list$, this.searchString$, (list, searchString) =>
            this.extractFilteredList(list, searchString)
        ),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error),
        searchString: this.select((state) => state.searchString)
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

    initFilters = this.updater((state) => ({
        ...state,
        ...pick(initialState, 'searchString')
    }));

    updateFilterSearchTerm = this.updater((state, searchTerm: string) => ({
        ...state,
        searchString: searchTerm
    }));

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

    private extractFilteredList(list: CompaniesListModel | null, searchString: string): CompaniesListModel | null {
        if (!list) {
            return null;
        }

        const filteredList = { ...list };

        const shouldFilterByTerm = searchString.length >= SEARCH_TERM_MIN_LEN;
        if (shouldFilterByTerm) {
            const fuse = new Fuse(filteredList.items, {
                keys: ['name', 'address', 'email'],
                shouldSort: false,
                threshold: 0.4
            });
            filteredList.items = fuse.search(searchString).map(({ refIndex }) => filteredList.items[refIndex]);
        }
        return filteredList;
    }
}
