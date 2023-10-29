import { exhaustMap, map, Observable, tap, withLatestFrom } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';
import { Store } from '@ngrx/store';

import { UserKindEnum } from '@app/common/models/users.model';
import { selectProfile } from '@app/common/stores/auth/auth.reducer';
import { CompaniesService } from '@app/companies/common/services/companies.service';
import { JobList } from '@app/jobs/common/models/job.model';
import { UsersList } from '@app/users/common/models/users-list.model';
import { AdminUsersService } from '@app/users/common/services/admin-users.service';

export interface DashboardState {
    jobs: JobList | null;
    users: UsersList | null;
    loading: boolean;
    error: boolean;
}

const initialState: DashboardState = {
    jobs: null,
    users: null,
    loading: false,
    error: false
};

@Injectable()
export class CompanyDashboardStore extends ComponentStore<DashboardState> {
    vm$ = this.select({
        jobs: this.select((state) => state.jobs),
        users: this.select((state) => state.users),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

    getJobs = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.dashboardLoading()),
            exhaustMap((id: number) =>
                this.companiesService.getJobsByCompany(id).pipe(
                    tapResponse(
                        (jobs) =>
                            this.patchState({
                                jobs,
                                loading: false
                            }),
                        () => this.dashboardLoadError
                    )
                )
            )
        )
    );

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
                                  (item) => item.kind === UserKindEnum.COMPANY && item.companyId === profile.companyId
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
        private readonly usersService: AdminUsersService,
        private readonly companiesService: CompaniesService
    ) {
        super(initialState);
    }
}
