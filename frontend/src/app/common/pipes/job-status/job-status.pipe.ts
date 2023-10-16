import { Pipe, PipeTransform } from '@angular/core';

import { JobStatusEnum } from '@app/common/models/jobs.model';

@Pipe({
    name: 'jobStatusLabel',
    standalone: true
})
export class JobStatusPipe implements PipeTransform {
    transform(status: JobStatusEnum): string {
        switch (status) {
            case JobStatusEnum.NOT_SEARCHING:
                return 'Currently not searching for job';
            case JobStatusEnum.OPEN_TO:
                return 'Open to opportunities';
            case JobStatusEnum.SEARCHING:
                return 'Searching for job';
            default:
                return 'N/A';
        }
    }
}
