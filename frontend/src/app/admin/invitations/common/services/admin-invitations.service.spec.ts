import { TestBed } from '@angular/core/testing';

import { AdminInvitationsService } from './admin-invitations.service';

describe('AdminInvitationsService', () => {
    let service: AdminInvitationsService;

    beforeEach(() => {
        TestBed.configureTestingModule({});
        service = TestBed.inject(AdminInvitationsService);
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });
});
