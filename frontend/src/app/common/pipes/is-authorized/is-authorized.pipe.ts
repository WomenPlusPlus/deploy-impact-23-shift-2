import { Pipe, PipeTransform } from '@angular/core';
import { Observable, of } from 'rxjs';
import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';

@Pipe({
    name: 'isAuthorized',
    standalone: true
})
export class IsAuthorizedPipe implements PipeTransform {

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    transform(isTrue: boolean, _kinds: UserKindEnum | UserKindEnum[], _roles: UserRoleEnum | UserRoleEnum[] = []): Observable<boolean> {
        return of(isTrue);
    }

}
