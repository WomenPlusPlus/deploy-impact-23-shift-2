import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateInviteComponent } from './create-invite.component';

describe('CreateInviteComponent', () => {
    let component: CreateInviteComponent;
    let fixture: ComponentFixture<CreateInviteComponent>;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [CreateInviteComponent]
        });
        fixture = TestBed.createComponent(CreateInviteComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
