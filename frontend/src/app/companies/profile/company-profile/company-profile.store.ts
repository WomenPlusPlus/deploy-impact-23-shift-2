import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { CompaniesService } from '@app/companies/common/services/companies.service';
import { JobList } from '@app/jobs/common/models/job.model';

import { CompanyProfileModel } from '../common/models/company-profile.model';

export interface CompanyProfileState {
    profile: CompanyProfileModel | null;
    jobs: JobList | null;
    loading: boolean;
    error: boolean;
}

const initialState: CompanyProfileState = {
    profile: null,
    jobs: null,
    loading: false,
    error: false
};

@Injectable()
export class CompanyProfileStore extends ComponentStore<CompanyProfileState> {
    vm$ = this.select({
        profile: this.select((state) => state.profile),
        jobs: this.select((state) => state.jobs),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

    getProfile = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.profileLoading()),
            exhaustMap((id: number) =>
                this.companiesService.getCompany(id).pipe(
                    tapResponse(
                        (profile) => this.profileLoadSuccess(profile),
                        () => {
                            this.profileLoadError();
                            this.router.navigate(['/companies']);
                            this.toast.error('Company with id ' + id + ' not found.');
                        }
                    )
                )
            )
        )
    );

    getJobs = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            exhaustMap((id: number) =>
                this.companiesService.getJobsByCompany(id).pipe(
                    tapResponse(
                        (jobs) => this.jobsLoadSuccess(jobs),
                        () => this.toast.error('Could not load job listings.')
                    )
                )
            )
        )
    );

    private profileLoading = this.updater(
        (state): CompanyProfileState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private profileLoadSuccess = this.updater(
        (state, profile: CompanyProfileModel): CompanyProfileState => ({
            ...state,
            profile,
            loading: false
        })
    );
    private profileLoadError = this.updater(
        (state): CompanyProfileState => ({
            ...state,
            error: true,
            loading: false
        })
    );
    private jobsLoadSuccess = this.updater(
        (state, jobs: JobList): CompanyProfileState => ({
            ...state,
            jobs
        })
    );

    constructor(
        private readonly companiesService: CompaniesService,
        private router: Router,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
