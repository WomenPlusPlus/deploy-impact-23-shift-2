import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faEllipsisV, faExternalLink, faEye } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import { AssociationProfileModel } from '@app/associations/common/models/association-profile.model';

import { AssociationsListStore } from '../associations-list.store';

@Component({
    selector: 'app-association-card',
    standalone: true,
    imports: [CommonModule, RouterModule, FontAwesomeModule],
    templateUrl: './association-card.component.html'
})
export class AssociationCardComponent {
    @Input() association!: AssociationProfileModel;
    @Input() deleting!: boolean;

    // TODO: check if user is Admin
    isUserAdmin = true;

    protected readonly faEye = faEye;
    protected readonly faExternalLink = faExternalLink;
    protected readonly faEllipsisV = faEllipsisV;

    constructor(private readonly associationsListStore: AssociationsListStore) {}

    onDelete(): void {
        this.associationsListStore.deleteItem(this.association.id);
    }

    get disableDeleteAction(): boolean {
        return this.deleting;
    }
}
