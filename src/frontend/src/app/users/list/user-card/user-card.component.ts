import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faEllipsisV, faExternalLink, faEye } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, HostBinding, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { UserKindEnum, UserStateEnum } from '@app/common/models/users.model';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';
import { UserAssociationRoleLabelPipe } from '@app/common/pipes/user-association-role-label/user-association-role-label.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { UserStateLabelPipe } from '@app/common/pipes/user-state-label/user-state-label.pipe';
import {
    DescriptionListComponent,
    DescriptionListItemComponent
} from '@app/ui/description-list/description-list.component';
import { UserBadgesComponent } from '@app/ui/user-badges/user-badges.component';
import { UsersCandidateItem, UsersItem, UsersListMode } from '@app/users/common/models/users-list.model';
import { UsersListStore } from '@app/users/list/users-list.store';

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
        UserStateLabelPipe,
        UserBadgesComponent,
        IsAuthorizedPipe
    ],
    templateUrl: './user-card.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserCardComponent {
    @Input() user!: UsersItem;
    @Input() mode!: UsersListMode;
    @Input() deleting!: boolean;

    @HostBinding('class.relative') classRelative = true;

    $candidate = (user: UsersItem): UsersCandidateItem => {
        return user as UsersCandidateItem;
    };

    get disableDeleteAction(): boolean {
        return this.user.state === UserStateEnum.DELETED || this.deleting;
    }

    protected readonly faEye = faEye;
    protected readonly faExternalLink = faExternalLink;
    protected readonly userKindEnum = UserKindEnum;
    protected readonly userStateEnum = UserStateEnum;

    constructor(private readonly usersListStore: UsersListStore) {}

    onDelete(): void {
        this.usersListStore.deleteItem(this.user.id);
    }

    onToggleMode(): void {
        this.mode = this.mode === 'short' ? 'detailed' : 'short';
    }

    protected readonly faEllipsisV = faEllipsisV;
}
