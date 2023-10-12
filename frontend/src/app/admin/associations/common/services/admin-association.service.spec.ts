import { TestBed } from '@angular/core/testing';

import { AdminAssociationService } from './admin-association.service';

describe('AdminAssociationService', () => {
    let service: AdminAssociationService;

    beforeEach(() => {
        TestBed.configureTestingModule({});
        service = TestBed.inject(AdminAssociationService);
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });
});
