import { Pipe, PipeTransform } from '@angular/core';

import { AssociationUserRoleEnum } from '@app/common/models/users.model';

@Pipe({
    name: 'userAssociationRoleLabel',
    standalone: true
})
export class UserAssociationRoleLabelPipe implements PipeTransform {
    transform(role: AssociationUserRoleEnum): string {
        switch (role) {
            case AssociationUserRoleEnum.ADMIN:
                return 'Admin';
            case AssociationUserRoleEnum.USER:
                return 'User';
            default:
                return 'N/A';
        }
    }
}
