import { HotToastService } from '@ngneat/hot-toast';
import { sanitize } from 'dompurify';
import { marked } from 'marked';
import { exhaustMap, Observable, tap } from 'rxjs';

import { HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { Job } from '@app/jobs/common/models/job.model';
import { JobsService } from '@app/jobs/common/services/jobs.service';

export interface JobDetailsState {
    job: Job | null;
    loading: boolean;
    error: boolean;
}

const initialState: JobDetailsState = {
    job: null,
    loading: false,
    error: false
};

@Injectable()
export class JobDetailsStore extends ComponentStore<JobDetailsState> {
    vm$ = this.select({
        job: this.select((state) => {
            const job = state.job;
            if (!job) {
                return job;
            }
            return {
                ...job,
                company: {
                    ...job.company,
                    mission: sanitize(marked.parse(job.company.mission)),
                    values: sanitize(marked.parse(job.company.values))
                },
                benefits: job.benefits && sanitize(marked.parse(job.benefits)),
                candidateOverview: sanitize(marked.parse(job.candidateOverview)),
                overview: sanitize(marked.parse(job.overview)),
                rolesAndResponsibility: sanitize(marked.parse(job.rolesAndResponsibility))
            };
        }),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error)
    });

    getDetails = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.jobLoading()),
            exhaustMap((id: number) =>
                this.jobsService.getJobDetails(id).pipe(
                    tapResponse(
                        (job) => this.jobLoadSuccess(job),
                        (error: HttpErrorResponse) => {
                            this.jobLoadError();
                            error.status === 404 && this.router.navigate(['/jobs']);
                            this.toast.error(`Job with id ${id} not found.`);
                        }
                    )
                )
            )
        )
    );

    private jobLoading = this.updater(
        (state): JobDetailsState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private jobLoadSuccess = this.updater(
        (state, job: Job): JobDetailsState => ({
            ...state,
            job,
            loading: false
        })
    );
    private jobLoadError = this.updater(
        (state): JobDetailsState => ({
            ...state,
            error: true,
            loading: false
        })
    );

    constructor(
        private readonly jobsService: JobsService,
        private router: Router,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
