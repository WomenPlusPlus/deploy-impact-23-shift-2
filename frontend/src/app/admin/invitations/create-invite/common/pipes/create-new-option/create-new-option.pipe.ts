import { Pipe, PipeTransform } from '@angular/core';

type Option = { id: number; name: string };

@Pipe({
    name: 'createNewOption',
    standalone: true
})
export class CreateNewOptionPipe implements PipeTransform {
    transform<T extends Option>(value: T[], flag: boolean): T[] {
        if (!flag) {
            return value;
        }

        return [{ id: 0, name: 'Add new on creation...' } as T].concat(value);
    }
}
