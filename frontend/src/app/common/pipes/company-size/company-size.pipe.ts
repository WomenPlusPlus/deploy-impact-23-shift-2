import { Pipe, PipeTransform } from '@angular/core';

import { CompanySizeEnum } from '@app/common/models/companies.model';

@Pipe({
    name: 'companySizeLabel',
    standalone: true
})
export class CompanySizePipe implements PipeTransform {
    transform(size: CompanySizeEnum): string {
        switch (size) {
            case CompanySizeEnum.ANY:
                return 'Any';
            case CompanySizeEnum.LARGE:
                return 'Large';
            case CompanySizeEnum.MEDIUM:
                return 'Medium';
            case CompanySizeEnum.SMALL:
                return 'Small';
            default:
                return 'Unknown';
        }
    }
}
