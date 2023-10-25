import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faArrowRight } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router, RouterLink } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { CompanyDashboardStore } from './company-dashboard.store';

@Component({
    selector: 'app-company-dashboard',
    standalone: true,
    imports: [CommonModule, RouterLink, ContentLoadingComponent, JobTypePipe, JobLocationTypePipe, FontAwesomeModule],
    providers: [provideComponentStore(CompanyDashboardStore)],
    templateUrl: './company-dashboard.component.html'
})
export class CompanyDashboardComponent implements OnInit {
    id?: number;
    readonly vm$ = this.companyDashboardStore.vm$;

    protected readonly faArrowRight = faArrowRight;

    constructor(
        private readonly companyDashboardStore: CompanyDashboardStore,
        private router: Router
    ) {}

    ngOnInit(): void {
        this.id = 1;
        if (this.id) {
            this.companyDashboardStore.getJobs(this.id);
            this.companyDashboardStore.getUsers(this.id);
        } else {
            this.router.navigate(['/']);
        }
    }
}
