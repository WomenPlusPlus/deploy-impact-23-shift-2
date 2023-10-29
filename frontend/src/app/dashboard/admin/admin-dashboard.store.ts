import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { JobList } from '@app/jobs/common/models/job.model';
import { UsersList } from '@app/users/common/models/users-list.model';

import { AssociationList } from '../common/models/associations.model';
import { CompanyList } from '../common/models/company.model';
import { AdminDashboardService } from '../common/services/admin.service';

export interface DashboardState {
    companies: CompanyList | null;
    jobs: JobList | null;
    associations: AssociationList | null;
    users: UsersList | null;
    loading: boolean;
    error: boolean;
}

const initialState: DashboardState = {
    companies: null,
    jobs: null,
    associations: null,
    users: null,
    loading: false,
    error: false
};

@Injectable()
export class AdminDashboardStore extends ComponentStore<DashboardState> {
    vm$ = this.select({
        companies: this.select((state) => state.companies),
        jobs: this.select((state) => state.jobs),
        associations: this.select((state) => state.associations),
        users: this.select((state) => state.users),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

    getUsers = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.dashboardLoading()),
            exhaustMap(() =>
                this.service.getUsers().pipe(
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

    getCompanies = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.dashboardLoading()),
            exhaustMap(() =>
                this.service.getCompanies().pipe(
                    tapResponse(
                        (companies) =>
                            this.patchState({
                                companies,
                                loading: false
                            }),
                        () => this.dashboardLoadError
                    )
                )
            )
        )
    );

    getJobs = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.dashboardLoading()),
            exhaustMap(() =>
                this.service.getJobs().pipe(
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

    getAssociations = this.effect((trigger$: Observable<void>) =>
        trigger$.pipe(
            tap(() => this.dashboardLoading()),
            exhaustMap(() =>
                this.service.getAssociations().pipe(
                    tapResponse(
                        (associations) =>
                            this.patchState({
                                associations,
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

    constructor(private readonly service: AdminDashboardService) {
        super(initialState);
    }
}
