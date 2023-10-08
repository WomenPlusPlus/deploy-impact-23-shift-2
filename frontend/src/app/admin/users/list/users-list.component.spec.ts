import { EMPTY } from 'rxjs';

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { provideComponentStore } from '@ngrx/component-store';
import { provideMockStore } from '@ngrx/store/testing';

import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';
import { UsersListStore } from '@app/admin/users/list/users-list.store';

import { UsersListComponent } from './users-list.component';

describe('UsersListComponent', () => {
    let component: UsersListComponent;
    let fixture: ComponentFixture<UsersListComponent>;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [UsersListComponent],
            providers: [
                { provide: AdminUsersService, useValue: { getList: () => EMPTY } },
                provideMockStore(),
                provideComponentStore(UsersListStore)
            ]
        });
        fixture = TestBed.createComponent(UsersListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
