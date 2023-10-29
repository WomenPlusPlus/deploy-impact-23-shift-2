import Fuse from 'fuse.js';
import { pick } from 'lodash';
import { exhaustMap, Observable, of, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { Store } from '@ngrx/store';

import { InvitationList } from '@app/common/models/invitation.model';
import { Profile } from '@app/common/models/profile.model';
import { profile$ } from '@app/common/utils/auth.util';
import { InvitationsService } from '@app/invitations/common/services/invitations.service';

export interface InvitationListState {
    list: InvitationList | null;
    loading: boolean;
    error: boolean;
    searchTerm: string;
    onlyMine: boolean;
}

const initialState: InvitationListState = {
    list: null,
    loading: false,
    error: false,
    searchTerm: '',
    onlyMine: false
};

const SEARCH_TERM_MIN_LEN = 3;

@Injectable()
export class InvitationListStore extends ComponentStore<InvitationListState> {
    private list$ = this.select((state) => state.list);

    searchTerm$ = this.select((state) => state.searchTerm);
    onlyMine$ = this.select((state) => state.onlyMine);
    vm$ = this.select({
        list: this.select(
            this.list$,
            this.searchTerm$,
            this.onlyMine$.pipe(switchMap((mine) => (mine ? profile$(this.store) : of(null)))),
            this.extractFilteredList.bind(this)
        ),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error),
        searchTerm: this.searchTerm$,
        onlyMine: this.onlyMine$
    });

    getList = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.getListLoading()),
            exhaustMap(() =>
                this.invitationsService.getList().pipe(
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
        ...pick(initialState, 'searchTerm', 'onlyMine')
    }));

    updateFilterOnlyMine = this.updater((state, onlyMine: boolean) => ({
        ...state,
        onlyMine
    }));

    updateFilterSearchTerm = this.updater((state, searchTerm: string) => ({
        ...state,
        searchTerm
    }));

    private getListLoading = this.updater(
        (state): InvitationListState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private getListLoadedSuccess = this.updater(
        (state, list: InvitationList): InvitationListState => ({
            ...state,
            list,
            loading: false
        })
    );
    private getListLoadedError = this.updater(
        (state): InvitationListState => ({
            ...state,
            error: true,
            loading: false
        })
    );

    constructor(
        private readonly store: Store,
        private readonly invitationsService: InvitationsService
    ) {
        super(initialState);
    }

    private extractFilteredList(
        list: InvitationList | null,
        searchTerm: string,
        profile: Profile | null
    ): InvitationList | null {
        if (!list) {
            return null;
        }

        const filteredList = { ...list };

        if (profile) {
            filteredList.items = filteredList.items.filter((item) => item.creatorId === profile.id);
        }

        const shouldFilterByTerm = searchTerm.length >= SEARCH_TERM_MIN_LEN;
        if (shouldFilterByTerm) {
            const fuse = new Fuse(filteredList.items, {
                keys: ['kind', 'role', 'email', 'state'],
                shouldSort: false,
                threshold: 0.4
            });
            filteredList.items = fuse.search(searchTerm).map(({ refIndex }) => filteredList.items[refIndex]);
        }
        return filteredList;
    }
}
