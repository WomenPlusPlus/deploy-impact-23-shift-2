import { HotToastService } from '@ngneat/hot-toast';
import Fuse from 'fuse.js';
import { pick } from 'lodash';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { UserKindEnum, UserStateEnum } from '@app/common/models/users.model';
import { UsersListMode, UsersListModel } from '@app/users/common/models/users-list.model';
import { AdminUsersService } from '@app/users/common/services/admin-users.service';

export interface UsersListState {
    list: UsersListModel | null;
    loading: boolean;
    error: boolean;
    deleting: boolean;
    filters: UsersListFiltersState;
    mode: UsersListMode;
}

export interface UsersListFiltersState {
    searchTerm: string;
    kinds: { [kind in UserKindEnum]: boolean };
}

const initialState: UsersListState = {
    list: null,
    loading: false,
    deleting: false,
    error: false,
    filters: {
        searchTerm: '',
        kinds: {
            [UserKindEnum.ADMIN]: false,
            [UserKindEnum.ASSOCIATION]: false,
            [UserKindEnum.CANDIDATE]: false,
            [UserKindEnum.COMPANY]: false
        }
    },
    mode: 'short'
};

const SEARCH_TERM_MIN_LEN = 3;

@Injectable()
export class UsersListStore extends ComponentStore<UsersListState> {
    private list$ = this.select((state) => state.list);

    mode$ = this.select((state) => state.mode);
    filters$ = this.select((state) => state.filters);
    filterKinds$ = this.select(this.filters$, (filters) => filters.kinds);
    filterSearchTerm$ = this.select(this.filters$, (filters) => filters.searchTerm);
    deleting$ = this.select((state) => state.deleting);
    vm$ = this.select({
        list: this.select(this.list$, this.filters$, (list, filters) => this.extractFilteredList(list, filters)),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error),
        filters: this.filters$,
        mode: this.mode$,
        deleting: this.deleting$
    });

    getList = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.getListLoading()),
            exhaustMap(() =>
                this.adminUsersService.getList().pipe(
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
                this.adminUsersService.deleteUser(id).pipe(
                    tapResponse(
                        () => {
                            const items = this.state().list?.items || [];
                            items.forEach((item) => item.id === id && (item.state = UserStateEnum.DELETED));
                            this.patchState({
                                deleting: false,
                                list: { items }
                            });
                        },
                        () => {
                            this.patchState({ deleting: false });
                            this.toast.error('Could not delete user! Please try again later or contact the support.');
                        }
                    )
                )
            )
        )
    );

    initFilters = this.updater((state) => ({
        ...state,
        ...pick(initialState, 'filters', 'mode')
    }));

    toggleMode = this.updater((state) => ({ ...state, mode: state.mode !== 'detailed' ? 'detailed' : 'short' }));

    updateFilterKind = this.updater((state, kind: UserKindEnum) => ({
        ...state,
        filters: {
            ...state.filters,
            kinds: { ...state.filters.kinds, [kind]: !state.filters.kinds[kind] }
        }
    }));
    updateFilterSearchTerm = this.updater((state, searchTerm: string) => ({
        ...state,
        filters: { ...state.filters, searchTerm }
    }));

    private getListLoading = this.updater(
        (state): UsersListState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private getListLoadedSuccess = this.updater(
        (state, list: UsersListModel): UsersListState => ({
            ...state,
            list,
            loading: false
        })
    );
    private getListLoadedError = this.updater(
        (state): UsersListState => ({
            ...state,
            error: true,
            loading: false
        })
    );

    constructor(
        private readonly adminUsersService: AdminUsersService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }

    private extractFilteredList(list: UsersListModel | null, filters: UsersListFiltersState): UsersListModel | null {
        if (!list) {
            return null;
        }

        const filteredList = { ...list };

        const shouldFilterByKinds = Object.keys(filters.kinds).some((kind) => filters.kinds[kind as UserKindEnum]);
        if (shouldFilterByKinds) {
            filteredList.items = filteredList.items.filter((item) => filters.kinds[item.kind]);
        }

        const shouldFilterByTerm = filters.searchTerm.length >= SEARCH_TERM_MIN_LEN;
        if (shouldFilterByTerm) {
            const fuse = new Fuse(filteredList.items, {
                keys: ['firstName', 'lastName', 'preferredName', 'email'],
                shouldSort: false,
                threshold: 0.4
            });
            filteredList.items = fuse.search(filters.searchTerm).map(({ refIndex }) => filteredList.items[refIndex]);
        }
        return filteredList;
    }
}
