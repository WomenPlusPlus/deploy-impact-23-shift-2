import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import environment from '@envs/environment';

import { Language, LocationCity } from '@app/common/models/location.model';

@Injectable({
    providedIn: 'root'
})
export class LocationsService {
    constructor(private readonly httpClient: HttpClient) {}

    getCities(): Observable<LocationCity[]> {
        return this.httpClient.get<LocationCity[]>(`${environment.API_BASE_URL}/locations/city`);
    }

    getLanguages(): Observable<Language[]> {
        return this.httpClient.get<Language[]>(`${environment.API_BASE_URL}/locations/language`);
    }
}
