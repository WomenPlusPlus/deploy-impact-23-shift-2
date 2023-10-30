import { Pipe, PipeTransform } from '@angular/core';

import { JobTypeEnum } from '@app/common/models/jobs.model';

@Pipe({
    name: 'jobTypeLabel',
    standalone: true
})
export class JobTypePipe implements PipeTransform {
    transform(type: JobTypeEnum): string {
        switch (type) {
            case JobTypeEnum.FULL_TIME:
                return 'Full-time';
            case JobTypeEnum.INTERNSHIP:
                return 'Internship';
            case JobTypeEnum.PART_TIME:
                return 'Part-time';
            case JobTypeEnum.TEMPORARY:
                return 'Temporary';
            default:
                return 'N/A';
        }
    }
}
