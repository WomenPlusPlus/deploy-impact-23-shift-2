import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { UsersListModel } from '@app/admin/users/common/models/users-list.model';
import { CreateUserResponse } from '@app/admin/users/form/common/models/create-user.model';
import { UserDetails } from '@app/common/models/users.model';
import { UserFormModel } from '@app/admin/users/form/common/models/user-form.model';
import { EditUserResponse } from '@app/admin/users/form/common/models/edit-user.model';

@Injectable({
    providedIn: 'root'
})
export class AdminUsersService {
    constructor(private readonly httpClient: HttpClient) {
    }

    getById(id: number): Observable<UserDetails> {
        return this.httpClient.get<UserDetails>(`${environment.API_BASE_URL}/api/v1/users/${id}`);
    }

    getList(): Observable<UsersListModel> {
        return this.httpClient
            .get<UsersListModel>(`${environment.API_BASE_URL}/api/v1/users`);
    }

    createUser(user: UserFormModel): Observable<CreateUserResponse> {
        return this.httpClient.post<CreateUserResponse>(`${environment.API_BASE_URL}/api/v1/users`, this.mapUserToFormData(user));
    }

    editUser(id: number, user: UserFormModel): Observable<EditUserResponse> {
        return this.httpClient.put<EditUserResponse>(`${environment.API_BASE_URL}/api/v1/users/${id}`, this.mapUserToFormData(user));
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
