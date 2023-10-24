import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { JobsListState, JobsListStore } from '@app/jobs/list/jobs-list.store';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

@Component({
    selector: 'app-jobs-list',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        JobTypePipe,
        JobLocationTypePipe,
        ContentErrorComponent,
        ContentLoadingComponent,
        FormsModule,
        LetDirective,
        UserKindLabelPipe
    ],
    providers: [JobsListStore],
    templateUrl: './jobs-list.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class JobsListComponent implements OnInit {
    vm$: Observable<JobsListState> = this.jobListStore.vm$;

    constructor(private readonly jobListStore: JobsListStore) {}

    ngOnInit(): void {
        this.loadData();
    }

    onMyJobsChange(state: boolean): void {
        this.jobListStore.updateFilterOnlyMine(state);
    }

    onSearchTermChange(term: string): void {
        this.jobListStore.updateFilterSearchTerm(term);
    }

    private loadData(): void {
        this.jobListStore.initFilters();
        this.jobListStore.getList();
    }
}
