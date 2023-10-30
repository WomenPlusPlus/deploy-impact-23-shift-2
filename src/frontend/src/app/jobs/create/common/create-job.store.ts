import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { JobsService } from '@app/jobs/common/services/jobs.service';
import { CreateJobFormModel } from '@app/jobs/create/common/models/create-job.model';

export interface CreateJobState {
    submitting: boolean;
    submitted: boolean;
}

const initialState: CreateJobState = {
    submitting: false,
    submitted: false
};

@Injectable()
export class CreateJobStore extends ComponentStore<CreateJobState> {
    vm$ = this.select({
        submitting: this.select((state) => state.submitting),
        submitted: this.select((state) => state.submitted)
    });

    submitForm = this.effect((trigger$: Observable<CreateJobFormModel>) =>
        trigger$.pipe(
            tap(() => this.patchState({ submitting: true, submitted: false })),
            exhaustMap((payload) =>
                this.jobsService.createJob(payload).pipe(
                    tapResponse(
                        () => this.patchState({ submitting: false, submitted: true }),
                        () => {
                            this.toast.error(
                                "Could not create job. Please try again later or contact the site's administrator."
                            );
                            this.patchState({ submitting: false, submitted: false });
                        }
                    )
                )
            )
        )
    );

    constructor(
        private readonly jobsService: JobsService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
