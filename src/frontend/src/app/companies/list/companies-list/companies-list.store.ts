import { HotToastService } from '@ngneat/hot-toast';
import Fuse from 'fuse.js';
import { pick } from 'lodash';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { CompaniesService } from '@app/companies/common/services/companies.service';
import { CompanyList } from '@app/companies/profile/common/models/company-profile.model';

export interface CompaniesListState {
    list: CompanyList | null;
    loading: boolean;
    error: boolean;
    deleting: boolean;
    searchString: string;
}

const initialState: CompaniesListState = {
    list: null,
    loading: false,
    error: false,
    deleting: false,
    searchString: ''
};

const SEARCH_TERM_MIN_LEN = 3;

@Injectable()
export class CompaniesListStore extends ComponentStore<CompaniesListState> {
    private list$ = this.select((state) => state.list);
    searchString$ = this.select((state) => state.searchString);
    deleting$ = this.select((state) => state.deleting);

    vm$ = this.select({
        list: this.select(this.list$, this.searchString$, (list, searchString) =>
            this.extractFilteredList(list, searchString)
        ),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error),
        deleting: this.deleting$,
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

    deleteItem = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.patchState({ deleting: true })),
            exhaustMap((id: number) =>
                this.companiesService.deleteCompany(id).pipe(
                    tapResponse(
                        () => {
                            const items = this.state().list?.items || [];
                            this.patchState({
                                deleting: false,
                                list: { items: items.filter((item) => item.id !== id) }
                            });
                        },
                        () => {
                            this.patchState({ deleting: false });
                            this.toast.error(
                                'Could not delete company! Please try again later or contact the support.'
                            );
                        }
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
        (state, list: CompanyList): CompaniesListState => ({
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

    constructor(
        private readonly companiesService: CompaniesService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }

    private extractFilteredList(list: CompanyList | null, searchString: string): CompanyList | null {
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
