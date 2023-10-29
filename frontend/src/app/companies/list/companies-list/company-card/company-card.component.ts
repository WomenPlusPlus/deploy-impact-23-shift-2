import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faEllipsisV } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import { UserKindEnum } from '@app/common/models/users.model';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';
import { UserAssociationRoleLabelPipe } from '@app/common/pipes/user-association-role-label/user-association-role-label.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { CompanyItem } from '@app/companies/profile/common/models/company-profile.model';

import { CompaniesListStore } from '../companies-list.store';

@Component({
    selector: 'app-company-card',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        FontAwesomeModule,
        IsAuthorizedPipe,
        UserAssociationRoleLabelPipe,
        UserCompanyRoleLabelPipe,
        UserKindLabelPipe
    ],
    templateUrl: './company-card.component.html'
})
export class CompanyCardComponent {
    @Input() company!: CompanyItem;
    @Input() deleting!: boolean;

    protected readonly faEllipsisV = faEllipsisV;

    protected readonly userKindEnum = UserKindEnum;

    constructor(private readonly companiesListStore: CompaniesListStore) {}

    onDelete(): void {
        this.companiesListStore.deleteItem(this.company.id);
    }

    get disableDeleteAction(): boolean {
        return this.deleting;
    }
}
