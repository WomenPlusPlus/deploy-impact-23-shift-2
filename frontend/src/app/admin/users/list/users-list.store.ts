import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { UserListModel } from '@app/admin/users/common/models/user-card.model';
import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';

export interface UsersListState {
    list: UserListModel | null;
    loading: boolean;
    error: boolean;
}

const initialState: UsersListState = {
    list: null,
    loading: false,
    error: false
};

@Injectable()
export class UsersListStore extends ComponentStore<UsersListState> {
    vm$ = this.select({
        list: this.select((state) => state.list),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
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

    private getListLoading = this.updater(
        (state): UsersListState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private getListLoadedSuccess = this.updater(
        (state, list: UserListModel): UsersListState => ({
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

    constructor(private readonly adminUsersService: AdminUsersService) {
        super(initialState);
    }
}
