import { Pipe, PipeTransform } from '@angular/core';

import { CompanyUserRoleEnum } from '@app/common/models/users.model';

@Pipe({
    name: 'userCompanyRoleLabel',
    standalone: true
})
export class UserCompanyRoleLabelPipe implements PipeTransform {
    transform(role: CompanyUserRoleEnum): string {
        switch (role) {
            case CompanyUserRoleEnum.ADMIN:
                return 'Admin';
            case CompanyUserRoleEnum.USER:
                return 'User';
            default:
                return 'N/A';
        }
    }
}
