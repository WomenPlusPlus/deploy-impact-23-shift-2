import { Observable, of } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { UsersListModel } from '@app/admin/users/common/models/users-list.model';
import { CreateUserFormModel } from '@app/admin/users/creation/common/models/create-user.model';
import { JobStatusEnum } from '@app/common/models/jobs.model';
import { UserRoleEnum, UserKindEnum, UserStateEnum } from '@app/common/models/users.model';

@Injectable({
    providedIn: 'root'
})
export class AdminUsersService {
    constructor(private readonly httpClient: HttpClient) {}

    getList(): Observable<UsersListModel> {
        // TODO: once the API is defined, delete this mock and use the endpoint instead.
        return of({
            items: [
                {
                    id: 0,
                    firstName: 'Test',
                    lastName: '123',
                    preferredName: 'Testing Admin',
                    imageUrl:
                        'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
                    email: 'test@test.com',
                    kind: UserKindEnum.ADMIN,
                    state: UserStateEnum.ACTIVE
                },
                {
                    id: 1,
                    firstName: 'João',
                    lastName: 'Rodrigues',
                    preferredName: 'John Cena',
                    imageUrl: 'https://cdn.wrestletalk.com/wp-content/uploads/2023/09/john-cena-september-2-d.jpg',
                    email: 'john-cena@mail.com',
                    kind: UserKindEnum.CANDIDATE,
                    state: UserStateEnum.DELETED,
                    phoneNumber: '999 000 555',
                    ratingSkill: 10,
                    jobStatus: JobStatusEnum.TEMPORARY,
                    hasCV: true,
                    hasVideo: false
                },
                {
                    id: 2,
                    firstName: 'Katy',
                    lastName: 'Perry',
                    imageUrl:
                        'https://www.koimoi.com/wp-content/new-galleries/2023/08/when-katy-perry-got-naughty-about-her-intimate-sx-life-deets-inside-001.jpg',
                    email: 'kat@perry.com',
                    kind: UserKindEnum.COMPANY,
                    state: UserStateEnum.ACTIVE,
                    role: UserRoleEnum.ADMIN
                },
                {
                    id: 3,
                    firstName: 'Tom',
                    lastName: 'Jerry',
                    preferredName: 'Tom&Jerry',
                    imageUrl: 'https://myinspiringthoughts.com/wp-content/uploads/2020/12/tomjerry1.jpg',
                    email: 'tom-jerry@movies.com',
                    kind: UserKindEnum.COMPANY,
                    state: UserStateEnum.ANONYMOUS,
                    role: UserRoleEnum.USER
                },
                {
                    id: 4,
                    firstName: 'Tom',
                    lastName: 'Cruise',
                    imageUrl:
                        'https://images.hindustantimes.com/img/2022/07/03/1600x900/Tom_Cruise_Top_Gun_Maverick_1656809669304_1656809705950.jpg',
                    email: 'tom-cruise@movies.com',
                    kind: UserKindEnum.ASSOCIATION,
                    state: UserStateEnum.ACTIVE,
                    role: UserRoleEnum.ADMIN
                },
                {
                    id: 5,
                    firstName: 'Céline',
                    lastName: 'Dion',
                    imageUrl: 'https://assets.medpagetoday.net/media/images/102xxx/102195.jpg?width=0.6',
                    email: 'dion@email.com',
                    kind: UserKindEnum.ASSOCIATION,
                    state: UserStateEnum.ACTIVE,
                    role: UserRoleEnum.USER
                }
            ]
        });
        /*return this.httpClient
            .get<UsersListModel>(`${environment.API_BASE_URL}/api/v1/users`);*/
    }

    createUser(user: CreateUserFormModel): Observable<{ id: number }> {
        const formData = new FormData();
        for (const key of Object.keys(user)) {
            const wrapper = user[key as keyof CreateUserFormModel];
            for (const key of Object.keys(wrapper)) {
                formData.append(key, wrapper[key as keyof typeof wrapper]);
            }
        }
        return this.httpClient.post<{ id: number }>(`${environment.API_BASE_URL}/api/v1/users`, formData);
    }

    deleteUser(id: number): Observable<void> {
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/users/${id}`);
    }
}
