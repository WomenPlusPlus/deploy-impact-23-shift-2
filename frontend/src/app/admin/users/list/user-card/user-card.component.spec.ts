import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { UserKindEnum, UserStateEnum } from '@app/common/models/users.model';

import { UserCardComponent } from './user-card.component';

describe('UserCardComponent', () => {
    let component: UserCardComponent;
    let fixture: ComponentFixture<UserCardComponent>;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [UserCardComponent, RouterTestingModule]
        });
        fixture = TestBed.createComponent(UserCardComponent);
        component = fixture.componentInstance;

        component.user = {
            id: 0,
            firstName: 'Test',
            lastName: '123',
            preferredName: 'Testing Admin',
            imageUrl:
                'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png',
            email: 'test@test.com',
            kind: UserKindEnum.ADMIN,
            state: UserStateEnum.ACTIVE
        };
        component.mode = 'detailed';
    });

    it('should create', () => {
        fixture.detectChanges();
        expect(component).toBeTruthy();
    });
});
