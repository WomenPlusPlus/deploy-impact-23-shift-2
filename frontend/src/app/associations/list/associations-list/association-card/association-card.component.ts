import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import { AssociationProfileModel } from '@app/associations/common/models/association-profile.model';

@Component({
    selector: 'app-association-card',
    standalone: true,
    imports: [CommonModule, RouterModule],
    templateUrl: './association-card.component.html'
})
export class AssociationCardComponent {
    @Input() association!: AssociationProfileModel;
}
