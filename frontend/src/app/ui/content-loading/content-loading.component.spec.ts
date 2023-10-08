import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ContentLoadingComponent } from './content-loading.component';

describe('ContentLoadingComponent', () => {
    let component: ContentLoadingComponent;
    let fixture: ComponentFixture<ContentLoadingComponent>;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [ContentLoadingComponent]
        });
        fixture = TestBed.createComponent(ContentLoadingComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
