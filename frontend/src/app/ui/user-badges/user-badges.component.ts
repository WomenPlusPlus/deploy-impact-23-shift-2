import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input } from '@angular/core';

import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';
import { UserAssociationRoleLabelPipe } from '@app/common/pipes/user-association-role-label/user-association-role-label.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';

@Component({
    selector: 'app-user-badges',
    standalone: true,
    imports: [CommonModule, UserAssociationRoleLabelPipe, UserCompanyRoleLabelPipe, UserKindLabelPipe],
    templateUrl: './user-badges.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserBadgesComponent {
    @Input() kind!: UserKindEnum;
    @Input() role?: UserRoleEnum;

    protected readonly userKindEnum = UserKindEnum;
}
