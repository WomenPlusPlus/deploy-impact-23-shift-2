import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faArrowRight } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { RouterLink, Router } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { AssociationDashboardStore } from './association-dashboard.store';

@Component({
    selector: 'app-association-dashboard',
    standalone: true,
    imports: [CommonModule, RouterLink, ContentLoadingComponent, FontAwesomeModule],
    providers: [provideComponentStore(AssociationDashboardStore)],
    templateUrl: './association-dashboard.component.html'
})
export class AssociationDashboardComponent implements OnInit {
    id?: number;
    readonly vm$ = this.associationDashboardStore.vm$;

    protected readonly faArrowRight = faArrowRight;

    constructor(
        private readonly associationDashboardStore: AssociationDashboardStore,
        private router: Router
    ) {}

    ngOnInit(): void {
        this.id = 1;
        if (this.id) {
            this.associationDashboardStore.getUsers(this.id);
        } else {
            this.router.navigate(['/']);
        }
    }
}
