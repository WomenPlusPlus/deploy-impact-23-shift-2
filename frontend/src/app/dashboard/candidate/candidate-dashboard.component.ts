import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faArrowRight } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { RouterLink, Router } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { CandidateDashboardStore } from './candidate-dashboard.store';

@Component({
    selector: 'app-candidate-dashboard',
    standalone: true,
    imports: [CommonModule, RouterLink, ContentLoadingComponent, JobLocationTypePipe, JobTypePipe, FontAwesomeModule],
    providers: [provideComponentStore(CandidateDashboardStore)],
    templateUrl: './candidate-dashboard.component.html'
})
export class CandidateDashboardComponent implements OnInit {
    id?: number;
    readonly vm$ = this.candidateDashboardStore.vm$;

    protected readonly faArrowRight = faArrowRight;

    constructor(
        private readonly candidateDashboardStore: CandidateDashboardStore,
        private router: Router
    ) {}

    ngOnInit(): void {
        this.id = 1;
        if (this.id) {
            this.candidateDashboardStore.getJobs();
        } else {
            this.router.navigate(['/']);
        }
    }
}
