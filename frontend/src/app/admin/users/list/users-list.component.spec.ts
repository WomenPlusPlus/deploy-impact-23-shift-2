import { EMPTY } from 'rxjs';

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminUsersService } from '@app/admin/users/common/services/admin-users.service';

import { UsersListComponent } from './users-list.component';

describe('UsersListComponent', () => {
    let component: UsersListComponent;
    let fixture: ComponentFixture<UsersListComponent>;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [UsersListComponent],
            providers: [{ provide: AdminUsersService, useValue: { getList: () => EMPTY } }]
        });
        fixture = TestBed.createComponent(UsersListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
