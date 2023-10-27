import { TestBed } from '@angular/core/testing';

import { AssociationProfileService } from './association-profile.service';

describe('AssociationProfileService', () => {
    let service: AssociationProfileService;

    beforeEach(() => {
        TestBed.configureTestingModule({});
        service = TestBed.inject(AssociationProfileService);
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });
});
