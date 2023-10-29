import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { ProfileSetup } from '@app/common/models/profile.model';
import { UserDetails } from '@app/common/models/users.model';
import { UsersList } from '@app/users/common/models/users-list.model';
import { CreateUserResponse } from '@app/users/form/common/models/create-user.model';
import { EditUserResponse } from '@app/users/form/common/models/edit-user.model';
import { UserFormModel } from '@app/users/form/common/models/user-form.model';

@Injectable({
    providedIn: 'root'
})
export class UsersService {
    constructor(private readonly httpClient: HttpClient) {}

    getById(id: number): Observable<UserDetails> {
        return this.httpClient.get<UserDetails>(`${environment.API_BASE_URL}/api/v1/users/${id}`);
    }

    getList(): Observable<UsersList> {
        return this.httpClient.get<UsersList>(`${environment.API_BASE_URL}/api/v1/users`);
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

    getSetupInfo(): Observable<ProfileSetup> {
        return this.httpClient.get<ProfileSetup>(`${environment.API_BASE_URL}/api/v1/setup`);
    }

    setupUser(user: UserFormModel): Observable<CreateUserResponse> {
        return this.httpClient.post<CreateUserResponse>(
            `${environment.API_BASE_URL}/api/v1/setup`,
            this.mapUserToFormData(user)
        );
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
