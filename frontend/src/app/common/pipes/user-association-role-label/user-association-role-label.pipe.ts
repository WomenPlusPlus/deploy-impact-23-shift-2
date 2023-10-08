import { Pipe, PipeTransform } from '@angular/core';

import { UserRoleEnum } from '@app/common/models/users.model';

@Pipe({
    name: 'userAssociationRoleLabel',
    standalone: true
})
export class UserAssociationRoleLabelPipe implements PipeTransform {
    transform(role: UserRoleEnum): string {
        switch (role) {
            case UserRoleEnum.ADMIN:
                return 'Admin';
            case UserRoleEnum.USER:
                return 'User';
            default:
                return 'N/A';
        }
    }
}
