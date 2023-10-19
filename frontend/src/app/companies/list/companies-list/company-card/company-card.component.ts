import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faEye, faExternalLink, faEllipsisV } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import { UserKindEnum } from '@app/common/models/users.model';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';
import { CompanyProfileModel } from '@app/companies/profile/common/models/company-profile.model';

import { CompaniesListStore } from '../companies-list.store';

@Component({
    selector: 'app-company-card',
    standalone: true,
    imports: [CommonModule, RouterModule, FontAwesomeModule, IsAuthorizedPipe],
    templateUrl: './company-card.component.html'
})
export class CompanyCardComponent {
    @Input() company!: CompanyProfileModel;
    @Input() deleting!: boolean;

    protected readonly faEye = faEye;
    protected readonly faExternalLink = faExternalLink;
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
