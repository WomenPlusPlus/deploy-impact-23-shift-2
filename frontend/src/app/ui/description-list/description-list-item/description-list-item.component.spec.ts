import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DescriptionListItemComponent } from './description-list-item.component';

describe('DescriptionListItemComponent', () => {
    let component: DescriptionListItemComponent;
    let fixture: ComponentFixture<DescriptionListItemComponent>;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [DescriptionListItemComponent]
        });
        fixture = TestBed.createComponent(DescriptionListItemComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
