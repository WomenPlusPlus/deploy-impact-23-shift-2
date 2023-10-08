import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faExternalLink, faEye } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, HostBinding, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import {
    UsersListAssociationModel,
    UsersListCandidateModel,
    UsersListCompanyModel,
    UsersListItemModel,
    UsersListMode
} from '@app/admin/users/common/models/users-list.model';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { UserKindEnum, UserStateEnum } from '@app/common/models/users.model';
import { UserAssociationRoleLabelPipe } from '@app/common/pipes/user-association-role-label/user-association-role-label.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { UserStateLabelPipe } from '@app/common/pipes/user-state-label/user-state-label.pipe';
import {
    DescriptionListComponent,
    DescriptionListItemComponent
} from '@app/ui/description-list/description-list.component';

@Component({
    selector: 'app-user-card',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        FontAwesomeModule,
        LetDirective,
        DescriptionListComponent,
        DescriptionListItemComponent,
        UserKindLabelPipe,
        UserAssociationRoleLabelPipe,
        UserCompanyRoleLabelPipe,
        UserStateLabelPipe
    ],
    templateUrl: './user-card.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserCardComponent {
    @Input() user!: UsersListItemModel;
    @Input() mode!: UsersListMode;

    @HostBinding('class.relative') classRelative = true;

    $candidate = (user: UsersListItemModel): UsersListCandidateModel => {
        return user as UsersListCandidateModel;
    };

    $company = (user: UsersListItemModel): UsersListCompanyModel => {
        return user as UsersListCompanyModel;
    };

    $association = (user: UsersListItemModel): UsersListAssociationModel => {
        return user as UsersListAssociationModel;
    };

    protected readonly faEye = faEye;
    protected readonly faExternalLink = faExternalLink;
    protected readonly userKindEnum = UserKindEnum;
    protected readonly userStateEnum = UserStateEnum;

    onToggleMode(): void {
        this.mode = this.mode === 'short' ? 'detailed' : 'short';
    }
}
