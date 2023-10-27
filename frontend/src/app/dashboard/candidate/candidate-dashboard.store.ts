import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { JobListings } from '../common/models/jobs.model';
import { AdminDashboardService } from '../common/services/admin.service';

export interface DashboardState {
    jobs: JobListings | null;
    loading: boolean;
    error: boolean;
}

const initialState: DashboardState = {
    jobs: null,
    loading: false,
    error: false
};

@Injectable()
export class CandidateDashboardStore extends ComponentStore<DashboardState> {
    vm$ = this.select({
        jobs: this.select((state) => state.jobs),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

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
