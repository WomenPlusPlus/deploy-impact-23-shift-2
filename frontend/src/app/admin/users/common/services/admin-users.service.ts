import { Observable, of, throwError } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { UsersListModel } from '@app/admin/users/common/models/users-list.model';
import { CreateUserResponse } from '@app/admin/users/form/common/models/create-user.model';
import { EditUserResponse } from '@app/admin/users/form/common/models/edit-user.model';
import { UserFormModel } from '@app/admin/users/form/common/models/user-form.model';
import { CompanySizeEnum } from '@app/common/models/companies.model';
import { JobLocationTypeEnum, JobStatusEnum, JobTypeEnum, WorkPermitEnum } from '@app/common/models/jobs.model';
import { UserDetails, UserKindEnum, UserRoleEnum, UserStateEnum } from '@app/common/models/users.model';

@Injectable({
    providedIn: 'root'
})
export class AdminUsersService {
    constructor(private readonly httpClient: HttpClient) {}

    getById(id: number): Observable<UserDetails> {
        // TODO: once the API is defined, delete this mock and use the endpoint instead.
        const list = [
            {
                id: 1,
                kind: UserKindEnum.ADMIN,
                firstName: 'joao',
                lastName: 'rodrigues',
                preferredName: 'João',
                email: 'joaordev+1@gmail.com',
                phoneNumber: '999000333',
                birthDate: '2023-10-15T00:00:00Z',
                photo: null,
                linkedInUrl: 'https://www.linkedin.com/in/jo%C3%A3o-rodrigues-84268613b/',
                githubUrl: 'https://github.com/jotar910',
                portfolioUrl: 'https://joaordev.vercel.app/'
            },
            {
                id: 2,
                kind: UserKindEnum.CANDIDATE,
                firstName: 'joao',
                lastName: 'rodrigues',
                preferredName: 'João',
                email: 'joaordev+2@gmail.com',
                phoneNumber: '999000333',
                birthDate: '2023-10-15T00:00:00Z',
                photo: null,
                linkedInUrl: 'https://www.linkedin.com/in/jo%C3%A3o-rodrigues-84268613b/',
                githubUrl: 'https://github.com/jotar910',
                portfolioUrl: 'https://joaordev.vercel.app/',
                candidateId: 1,
                yearsOfExperience: 5,
                jobStatus: JobStatusEnum.OPEN_TO,
                seekJobType: JobTypeEnum.PART_TIME,
                seekCompanySize: CompanySizeEnum.SMALL,
                seekLocations: [
                    {
                        id: 32832,
                        name: 'Coimbra'
                    }
                ],
                seekLocationType: JobLocationTypeEnum.HYBRID,
                seekSalary: 1000000,
                seekValues: 'Whatever',
                workPermit: WorkPermitEnum.WORK_VISA,
                noticePeriod: 4,
                spokenLanguages: [
                    {
                        language: {
                            id: 1,
                            name: 'English',
                            shortName: 'EN'
                        },
                        level: 1
                    },
                    {
                        language: {
                            id: 8,
                            name: 'Portuguese',
                            shortName: 'PT'
                        },
                        level: 5
                    }
                ],
                skills: [
                    {
                        name: 'Frontend',
                        years: 5
                    },
                    {
                        name: 'Backend',
                        years: 2
                    }
                ],
                cv: {
                    name: 'CV_JoaoRodrigues.pdf',
                    url: 'https://storage.googleapis.com/app-shift2-bucket/1/cv/CV_JoaoRodrigues.pdf?Expires=1697793680&GoogleAccessId=app-shift2-bucket-sa%40shift-2-400920.iam.gserviceaccount.com&Signature=uzkDAew5a7pu5bloglu5GwooJc8gcD%2BxgcUOM%2BaGFC22InZI41P91P40w26H8z8pahtFHMk%2BN8kpftzVOVLyhLy7qFSxeE1hzW6YJuZyf6NHipw3Eh2rJHbg7gVrgEU6yHZ1MIeAHzsn5edm0fHmuQcQ5laUfhsOCyUfUeg3qOKbpfK%2Fy0OuB1ccxBZp9vCm8ak0MKtMe7Pmr%2B5E1XHksEmCKrwPMyeiKAhVxVSN9h79Vzdqg3Tfdyt97jVgwj2YoKMeLe1aiE5iv8PVXfoUyQxszI%2BFJGzC%2FtzKl%2BcbMh7DtqytgycHDeDFSHXCiWZdHQsNr3tGvkzYl643z0PLQQ%3D%3D'
                },
                attachments: [
                    {
                        name: 'empty.pdf',
                        url: 'https://storage.googleapis.com/app-shift2-bucket/1/attachments/empty.pdf?Expires=1697793680&GoogleAccessId=app-shift2-bucket-sa%40shift-2-400920.iam.gserviceaccount.com&Signature=B0JxyAiO6exD7UAt4m1Fc%2BQB6RckcScPbzkIE3jcs3%2B6sFGxVuUA9Rz%2F9mUVcJsFAmwM1orgZ1K0w%2Fsiq7%2B7L4qXlDLbJSWdj3VFGGj0fis0CTxLQeFclavA2Xv95ocJhNixh%2FeqeThQveIn7uR2v421V5K0Fj8Zv2%2BFsCglKFzUbwjZSipO2swbFLvXWR41bz9s3%2F3bJ40qyk0p5hA37BSUEmfHrFyX1NaWn7pkgYQ%2BT6d7ccpro8Bp6AMh%2FqighYi%2F7KpaqwSBSScQE1qwwokalZnd67apf8jlURPlBQBOeUGQQRxM7veQwvQHiIjDOjlqpvqeh0L5%2Bl9HxuPHYw%3D%3D'
                    },
                    {
                        name: 'empty.xlsx',
                        url: 'https://storage.googleapis.com/app-shift2-bucket/1/attachments/empty.xlsx?Expires=1697793680&GoogleAccessId=app-shift2-bucket-sa%40shift-2-400920.iam.gserviceaccount.com&Signature=ZdAQ4KW%2FEw%2FOMZlG9n9ofFRxUb6d1OI%2FEeKPmY2Ciua5A1o9XYDQgqBxCuunAiq3qmoIzGE35dnQ6KdZRbboSGn6uywCppHT%2B3Vwwa3JKqXUFrotAmnwkD9uvuu1h2vdTVVSz1JihFzE0OIjKr0VwzM9%2BpeWmDwkbIGqN0LDLpI3iEQESKXqXfwBpP2NM3LotoGDOkELopRgKxeSlP0198U92413R5hucz32wEL%2Facb6D2hs6%2BniuwbOf%2FNZVVuTTG4XmXAiSndwlFojLPTXdepcEoAtkvo66lIuF6fyhk6pvfMNI9vAct07mjlaxtRdnoF213j0NUV8xTPFycGJbg%3D%3D'
                    }
                ],
                video: {
                    name: 'Screen Recording 2023-10-08 at 13.26.53.mov',
                    url: 'https://storage.googleapis.com/app-shift2-bucket/1/video/Screen%20Recording%202023-10-08%20at%2013.26.53.mov?Expires=1697794711&GoogleAccessId=app-shift2-bucket-sa%40shift-2-400920.iam.gserviceaccount.com&Signature=UjYsjOGyn3V0dj68Fob1gycHMIflXYYzTLg6aDya9l3CEOhktUz7lSx6YBCaLp91W%2FM%2FDRVS4%2FHl2GTTyDb6bNVkqf2vAQG5yKcZbqWcGlHq768xUiLkAPPd9cwCpTHb%2Bm7ibay%2FfMjgKQJHvOSAHJnaW8jDu27%2BrCSQASCSYvVvsq1aMuyg%2B0K5GOa0hxKafxOZpvnzX3qeXr0ExYzMrLVDSceqBYS4A2WNId9AzbigJfhAz04AjmIqtpwefuwM0a8Fefetd70zps3ZRxTXx%2FINRYvCWGzNg%2F7kW%2FVylUGRPN4SbeSC8bFygPQGNzSI27D%2FQhB%2FOOh2Tt%2BOjGa8Tg%3D%3D'
                },
                educationHistory: [
                    {
                        title: 'Master',
                        description: 'Master in Software Engineering',
                        entity: 'ISEC',
                        fromDate: '2023-10-08T00:00:00Z',
                        toDate: '2023-10-15T00:00:00Z'
                    }
                ],
                employmentHistory: [
                    {
                        title: 'Software Engineer',
                        description: 'Fullstack developer',
                        company: 'Container xChange',
                        fromDate: '2023-10-15T00:00:00Z',
                        toDate: null
                    }
                ]
            },
            {
                id: 3,
                kind: UserKindEnum.ASSOCIATION,
                firstName: 'joao',
                lastName: 'rodrigues',
                preferredName: 'João',
                email: 'joaordev+3@gmail.com',
                phoneNumber: '999000333',
                birthDate: '2023-10-15T00:00:00Z',
                photo: null,
                linkedInUrl: 'https://www.linkedin.com/in/jo%C3%A3o-rodrigues-84268613b/',
                githubUrl: 'https://github.com/jotar910',
                portfolioUrl: 'https://joaordev.vercel.app/',
                associationUserId: 1,
                associationId: 1,
                role: UserRoleEnum.ADMIN
            },
            {
                id: 4,
                kind: UserKindEnum.COMPANY,
                firstName: 'joao',
                lastName: 'rodrigues',
                preferredName: 'João',
                email: 'joaordev+4@gmail.com',
                phoneNumber: '999000333',
                birthDate: '2023-10-15T00:00:00Z',
                photo: null,
                linkedInUrl: 'https://www.linkedin.com/in/jo%C3%A3o-rodrigues-84268613b/',
                githubUrl: 'https://github.com/jotar910',
                portfolioUrl: 'https://joaordev.vercel.app/',
                companyUserId: 1,
                companyId: 1,
                role: UserRoleEnum.USER
            }
        ];
        const item = list.find((item) => item.id === id);
        return item ? of(item) : throwError(() => ({ status: 404 }));
        // return this.httpClient.get<UserDetails>(`http://localhost:8080/api/v1/users/${id}`);
    }

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
                    jobStatus: JobStatusEnum.SEARCHING,
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

    createUser(user: UserFormModel): Observable<CreateUserResponse> {
        return this.httpClient.post<CreateUserResponse>(
            `${environment.API_BASE_URL}/api/v1/users`,
            this.mapUserToFormData(user)
        );
    }

    editUser(id: number, user: UserFormModel): Observable<EditUserResponse> {
        return this.httpClient.put<EditUserResponse>(
            `${environment.API_BASE_URL}/api/v1/users/${id}`,
            this.mapUserToFormData(user)
        );
    }

    deleteUser(id: number): Observable<void> {
        return this.httpClient.delete<void>(`${environment.API_BASE_URL}/api/v1/users/${id}`);
    }

    private mapUserToFormData(user: UserFormModel): FormData {
        const formData = new FormData();
        for (const key of Object.keys(user)) {
            const wrapper = user[key as keyof UserFormModel];
            if (typeof wrapper !== 'object') {
                formData.append(key, wrapper);
                continue;
            }
            for (const key of Object.keys(wrapper)) {
                const value: any = wrapper[key as keyof typeof wrapper];
                if (!value) {
                    formData.append(key, value);
                    continue;
                }
                if (value instanceof Date) {
                    formData.append(key, value.toISOString());
                    continue;
                }
                if (typeof value !== 'object' || value instanceof File) {
                    formData.append(key, value);
                    continue;
                }
                if (Array.isArray(value) && value[0] instanceof File) {
                    for (const v of value) {
                        formData.append(key, v);
                    }
                    continue;
                }
                if (key === 'photo' || key === 'cv' || key === 'video' || key === 'attachments') {
                    continue;
                }
                try {
                    formData.append(key, JSON.stringify(value));
                } catch (e) {
                    console.error(e);
                }
            }
        }
        return formData;
    }
}
