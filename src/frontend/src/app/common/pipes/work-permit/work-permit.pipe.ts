import { Pipe, PipeTransform } from '@angular/core';

import { WorkPermitEnum } from '@app/common/models/jobs.model';

@Pipe({
    name: 'workPermitLabel',
    standalone: true
})
export class WorkPermitPipe implements PipeTransform {
    transform(workPermit: WorkPermitEnum): string {
        switch (workPermit) {
            case WorkPermitEnum.CITIZEN:
                return 'Citizen';
            case WorkPermitEnum.PERMANENT_RESIDENT:
                return 'Permanent resident';
            case WorkPermitEnum.WORK_VISA:
                return 'Work visa';
            case WorkPermitEnum.STUDENT_VISA:
                return 'Student visa';
            case WorkPermitEnum.TEMPORARY_RESIDENT:
                return 'Temporary resident';
            case WorkPermitEnum.NO_WORK_PERMIT:
                return 'No work permit';
            case WorkPermitEnum.OTHER:
                return 'Other';
            default:
                return 'Unknown';
        }
    }
}
