import Fuse from 'fuse.js';
import { pick } from 'lodash';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { JobList } from '@app/jobs/common/models/job.model';
import { JobsService } from '@app/jobs/common/services/jobs.service';

export interface JobsListState {
    list: JobList | null;
    loading: boolean;
    error: boolean;
    searchTerm: string;
}

const initialState: JobsListState = {
    list: null,
    loading: false,
    error: false,
    searchTerm: ''
};

const SEARCH_TERM_MIN_LEN = 3;

@Injectable()
export class JobsListStore extends ComponentStore<JobsListState> {
    private list$ = this.select((state) => state.list);

    searchTerm$ = this.select((state) => state.searchTerm);
    vm$ = this.select({
        list: this.select(this.list$, this.searchTerm$, (list, searchTerm) =>
            this.extractFilteredList(list, searchTerm)
        ),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error),
        searchTerm: this.searchTerm$
    });

    getList = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.getListLoading()),
            exhaustMap(() =>
                this.jobsService.getList().pipe(
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
        ...pick(initialState, 'searchTerm')
    }));

    updateFilterSearchTerm = this.updater((state, searchTerm: string) => ({
        ...state,
        searchTerm
    }));

    private getListLoading = this.updater(
        (state): JobsListState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private getListLoadedSuccess = this.updater(
        (state, list: JobList): JobsListState => ({
            ...state,
            list,
            loading: false
        })
    );
    private getListLoadedError = this.updater(
        (state): JobsListState => ({
            ...state,
            error: true,
            loading: false
        })
    );

    constructor(private readonly jobsService: JobsService) {
        super(initialState);
    }

    private extractFilteredList(list: JobList | null, searchTerm: string): JobList | null {
        if (!list) {
            return null;
        }

        const filteredList = { ...list };

        const shouldFilterByTerm = searchTerm.length >= SEARCH_TERM_MIN_LEN;
        if (shouldFilterByTerm) {
            const fuse = new Fuse(filteredList.items, {
                keys: [
                    'title',
                    'company.name',
                    'creator.name',
                    'creator.email',
                    'location.city.name',
                    'location.type',
                    'creationDate'
                ],
                shouldSort: false,
                threshold: 0.4
            });
            filteredList.items = fuse.search(searchTerm).map(({ refIndex }) => filteredList.items[refIndex]);
        }
        return filteredList;
    }
}
