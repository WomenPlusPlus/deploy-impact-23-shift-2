import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { provideMockStore } from '@ngrx/store/testing';

import { AppStore } from '@app/app.store';

import { AppComponent } from './app.component';

describe('AppComponent', () => {
    beforeEach(() =>
        TestBed.configureTestingModule({
            imports: [AppComponent, RouterTestingModule],
            providers: [provideMockStore(), AppStore]
        })
    );

    it('should create the app', () => {
        const fixture = TestBed.createComponent(AppComponent);
        const app = fixture.componentInstance;
        expect(app).toBeTruthy();
    });
});
