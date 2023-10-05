import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

import { AdminUsersService } from './admin-users.service';

describe('AdminUsersService', () => {
    let service: AdminUsersService;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule]
        });
        service = TestBed.inject(AdminUsersService);
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });
});
