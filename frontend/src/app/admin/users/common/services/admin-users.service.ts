import { map, Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { UserListItemModel, UserListModel } from '../models/user-card.model';

@Injectable({
    providedIn: 'root'
})
export class AdminUsersService {
    constructor(private readonly httpClient: HttpClient) {}

    getList(): Observable<UserListModel> {
        return this.httpClient
            .get<UserListItemModel>(`${environment.API_BASE_URL}/`)
            .pipe(map((singleItem) => ({ items: [singleItem] })));
    }
}
