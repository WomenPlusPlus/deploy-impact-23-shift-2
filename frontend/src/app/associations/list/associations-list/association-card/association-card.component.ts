import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faEllipsisV, faExternalLink, faEye } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import { AssociationProfileModel } from '@app/associations/common/models/association-profile.model';
import { UserKindEnum } from '@app/common/models/users.model';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';

import { AssociationsListStore } from '../associations-list.store';

@Component({
    selector: 'app-association-card',
    standalone: true,
    templateUrl: './association-card.component.html',
    imports: [CommonModule, RouterModule, FontAwesomeModule, IsAuthorizedPipe]
})
export class AssociationCardComponent {
    @Input() association!: AssociationProfileModel;
    @Input() deleting!: boolean;

    protected readonly faEye = faEye;
    protected readonly faExternalLink = faExternalLink;
    protected readonly faEllipsisV = faEllipsisV;

    protected readonly userKindEnum = UserKindEnum;

    constructor(private readonly associationsListStore: AssociationsListStore) {}

    onDelete(): void {
        this.associationsListStore.deleteItem(this.association.id);
    }

    get disableDeleteAction(): boolean {
        return this.deleting;
    }
}
