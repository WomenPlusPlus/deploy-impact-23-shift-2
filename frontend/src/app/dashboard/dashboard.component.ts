import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';

import { UserKindEnum } from '@app/common/models/users.model';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';

import { AdminDashboardComponent } from './admin/admin-dashboard.component';
import { AssociationDashboardComponent } from './association/association-dashboard/association-dashboard.component';
import { CompanyDashboardComponent } from './company/company-dashboard/company-dashboard.component';

@Component({
    selector: 'app-dashboard',
    standalone: true,
    imports: [
        CommonModule,
        IsAuthorizedPipe,
        CompanyDashboardComponent,
        AssociationDashboardComponent,
        AdminDashboardComponent
    ],
    templateUrl: './dashboard.component.html'
})
export class DashboardComponent {
    protected readonly userKindEnum = UserKindEnum;
}
