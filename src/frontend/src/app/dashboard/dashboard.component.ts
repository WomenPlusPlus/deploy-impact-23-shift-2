import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';

import { Store } from '@ngrx/store';

import { Profile } from '@app/common/models/profile.model';
import { UserKindEnum } from '@app/common/models/users.model';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';
import { selectProfile } from '@app/common/stores/auth/auth.reducer';

import { AdminDashboardComponent } from './admin/admin-dashboard.component';
import { AssociationDashboardComponent } from './association/association-dashboard/association-dashboard.component';
import { CandidateDashboardComponent } from './candidate/candidate-dashboard.component';
import { CompanyDashboardComponent } from './company/company-dashboard/company-dashboard.component';

@Component({
    selector: 'app-dashboard',
    standalone: true,
    imports: [
        CommonModule,
        IsAuthorizedPipe,
        CompanyDashboardComponent,
        AssociationDashboardComponent,
        AdminDashboardComponent,
        CandidateDashboardComponent
    ],
    templateUrl: './dashboard.component.html'
})
export class DashboardComponent implements OnInit {
    protected readonly userKindEnum = UserKindEnum;

    profile$!: Observable<Profile | null>;

    constructor(private readonly store: Store) {}

    ngOnInit(): void {
        this.initSubscriptions();
    }

    private initSubscriptions(): void {
        this.profile$ = this.store.select(selectProfile);
    }
}
