import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { UsersList } from '@app/dashboard/common/models/users.model';
import { CompanyDasboardService } from '@app/dashboard/common/services/company.service';
import { JobList } from '@app/jobs/common/models/job.model';

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
                this.service.getJobsByCompany(id).pipe(
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
                this.service.getUsersByCompany().pipe(
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

    constructor(private readonly service: CompanyDasboardService) {
        super(initialState);
    }
}
