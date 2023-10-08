import { EMPTY } from 'rxjs';

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { provideComponentStore } from '@ngrx/component-store';
import { provideMockStore } from '@ngrx/store/testing';

import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';
import { UsersListStore } from '@app/admin/users/list/users-list.store';

import { UsersFiltersComponent } from './users-filters.component';

describe('UsersFiltersComponent', () => {
    let component: UsersFiltersComponent;
    let fixture: ComponentFixture<UsersFiltersComponent>;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [UsersFiltersComponent],
            providers: [
                { provide: AdminUsersService, useValue: { getList: () => EMPTY } },
                provideMockStore(),
                provideComponentStore(UsersListStore)
            ]
        });
        fixture = TestBed.createComponent(UsersFiltersComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
