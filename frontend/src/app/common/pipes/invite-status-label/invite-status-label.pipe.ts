import { Pipe, PipeTransform } from '@angular/core';

import { InviteStateEnum } from '@app/common/models/invitation.model';

@Pipe({
    name: 'inviteStatusLabel',
    standalone: true
})
export class InviteStatusLabelPipe implements PipeTransform {
    transform(state: InviteStateEnum, expire?: Date | string): string {
        switch (state) {
            case InviteStateEnum.CREATED:
                return 'Created';
            case InviteStateEnum.PENDING:
                if (expire && new Date(expire) < new Date()) {
                    return 'Expired';
                }
                return 'Pending';
            case InviteStateEnum.ERROR:
                return 'Failed';
            case InviteStateEnum.ACCEPTED:
                return 'Accepted';
            case InviteStateEnum.CANCELLED:
                return 'Cancelled';
            default:
                return 'N/A';
        }
    }
}
