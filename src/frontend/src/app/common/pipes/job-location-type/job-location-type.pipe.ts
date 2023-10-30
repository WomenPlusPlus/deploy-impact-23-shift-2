import { Pipe, PipeTransform } from '@angular/core';

import { JobLocationTypeEnum } from '@app/common/models/jobs.model';

@Pipe({
    name: 'jobLocationTypeLabel',
    standalone: true
})
export class JobLocationTypePipe implements PipeTransform {
    transform(locationType: JobLocationTypeEnum): string {
        switch (locationType) {
            case JobLocationTypeEnum.ON_SITE:
                return 'On-site';
            case JobLocationTypeEnum.HYBRID:
                return 'Hybrid';
            case JobLocationTypeEnum.REMOTE:
                return 'Remote';
            default:
                return 'Unknown';
        }
    }
}
