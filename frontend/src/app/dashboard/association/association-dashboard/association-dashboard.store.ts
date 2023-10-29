import { exhaustMap, map, Observable, tap, withLatestFrom } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { Store } from '@ngrx/store';

import { UserKindEnum } from '@app/common/models/users.model';
import { selectProfile } from '@app/common/stores/auth/auth.reducer';
import { UsersList } from '@app/users/common/models/users-list.model';
import { UsersService } from '@app/users/common/services/users.service';

export interface DashboardState {
    users: UsersList | null;
    loading: boolean;
    error: boolean;
}

const initialState: DashboardState = {
    users: null,
    loading: false,
    error: false
};

@Injectable()
export class AssociationDashboardStore extends ComponentStore<DashboardState> {
    vm$ = this.select({
        users: this.select((state) => state.users),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

    getUsers = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.dashboardLoading()),
            exhaustMap(() =>
                this.usersService.getList().pipe(
                    withLatestFrom(this.store.select(selectProfile)),
                    map(([users, profile]) => ({
                        items: !profile
                            ? []
                            : users.items.filter(
                                  (item) =>
                                      item.kind === UserKindEnum.ASSOCIATION &&
                                      item.associationId === profile.associationId
                              )
                    })),
                    tapResponse(
                        (users) =>
                            this.patchState({
                                users,
                                loading: false
                            }),
                        () => this.dashboardLoadError
                    )
                )
            )
        )
    );

    private dashboardLoading = this.updater(
        (state): DashboardState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private dashboardLoadError = this.updater(
        (state): DashboardState => ({
            ...state,
            error: true,
            loading: false
        })
    );

    constructor(
        private readonly store: Store,
        private readonly usersService: UsersService
    ) {
        super(initialState);
    }
}
