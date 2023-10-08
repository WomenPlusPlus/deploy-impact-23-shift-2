import { Pipe, PipeTransform } from '@angular/core';

import { UserStateEnum } from '@app/common/models/users.model';

@Pipe({
    name: 'userStateLabel',
    standalone: true
})
export class UserStateLabelPipe implements PipeTransform {
    transform(value: UserStateEnum): string {
        switch (value) {
            case UserStateEnum.ACTIVE:
                return 'Active';
            case UserStateEnum.ANONYMOUS:
                return 'Anonymous';
            case UserStateEnum.DELETED:
                return 'Deleted';
            default:
                return 'Unknown';
        }
    }
}
