import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { UsersList } from '@app/dashboard/common/models/users.model';
import { AssociationDasboardService } from '@app/dashboard/common/services/association.service';

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
                this.service.getUsersByAssociation().pipe(
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

    constructor(private readonly service: AssociationDasboardService) {
        super(initialState);
    }
}
