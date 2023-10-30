import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faArrowRight } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, OnInit, Input } from '@angular/core';
import { Router, RouterLink } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { UserKindEnum, UserStateEnum } from '@app/common/models/users.model';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';
import { UserBadgesComponent } from '@app/ui/user-badges/user-badges.component';

import { AdminDashboardStore } from './admin-dashboard.store';

@Component({
    selector: 'app-admin-dashboard',
    standalone: true,
    imports: [CommonModule, RouterLink, ContentLoadingComponent, UserBadgesComponent, FontAwesomeModule],
    providers: [provideComponentStore(AdminDashboardStore)],
    templateUrl: './admin-dashboard.component.html'
})
export class AdminDashboardComponent implements OnInit {
    @Input() id!: number;
    readonly vm$ = this.adminDashboardStore.vm$;

    protected readonly faArrowRight = faArrowRight;
    protected readonly userKindEnum = UserKindEnum;
    protected readonly userStateEnum = UserStateEnum;

    constructor(
        private readonly adminDashboardStore: AdminDashboardStore,
        private router: Router
    ) {}

    ngOnInit(): void {
        if (this.id) {
            this.adminDashboardStore.getUsers();
            this.adminDashboardStore.getAssociations();
            this.adminDashboardStore.getCompanies();
            this.adminDashboardStore.getJobs();
        } else {
            this.router.navigate(['/']);
        }
    }
}
