import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { JobDetailsComponent } from '@app/jobs/details/job-details.component';
import { SearchJobsState, SearchJobsStore } from '@app/jobs/search/search-jobs.store';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

@Component({
    selector: 'app-search-jobs',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        FormsModule,
        ContentErrorComponent,
        ContentLoadingComponent,
        JobLocationTypePipe,
        JobTypePipe,
        LetDirective,
        ReactiveFormsModule,
        JobDetailsComponent
    ],
    templateUrl: './search-jobs.component.html',
    providers: [SearchJobsStore],
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SearchJobsComponent implements OnInit {
    vm$: Observable<SearchJobsState> = this.searchJobsStore.vm$;

    selectedId: number | null = null;

    constructor(
        private readonly router: Router,
        private readonly searchJobsStore: SearchJobsStore
    ) {}

    ngOnInit(): void {
        this.loadData();
    }

    onSearchTermChange(term: string): void {
        this.searchJobsStore.updateFilterSearchTerm(term);
    }

    onJobSelect(id: number) {
        if (window.innerWidth < 1024) {
            // lg media query
            this.router.navigate(['/jobs', id]);
            return;
        }
        this.selectedId = this.selectedId === id ? null : id;
    }

    private loadData(): void {
        this.searchJobsStore.initFilters();
        this.searchJobsStore.getList();
    }
}
