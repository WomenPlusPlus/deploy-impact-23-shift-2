import { Pipe, PipeTransform } from '@angular/core';

import { UserKindEnum } from '@app/common/models/users.model';

@Pipe({
    name: 'userKindLabel',
    standalone: true
})
export class UserKindLabelPipe implements PipeTransform {
    transform(kind: UserKindEnum): string {
        switch (kind) {
            case UserKindEnum.ADMIN:
                return 'Administrator';
            case UserKindEnum.ASSOCIATION:
                return 'Association';
            case UserKindEnum.CANDIDATE:
                return 'Candidate';
            case UserKindEnum.COMPANY:
                return 'Company';
            default:
                return 'N/A';
        }
    }
}
