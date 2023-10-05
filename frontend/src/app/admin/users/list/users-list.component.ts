import { Observable, share } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';

import { UserListModel } from '../common/models/user-card.model';
import { AdminUsersService } from '../common/services/admin-users.service';
import { UserCardComponent } from './user-card/user-card.component';

@Component({
    selector: 'app-users-list',
    standalone: true,
    imports: [CommonModule, UserCardComponent],
    templateUrl: './users-list.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UsersListComponent implements OnInit {
    list$!: Observable<UserListModel>;

    constructor(private readonly adminUsersService: AdminUsersService) {}

    ngOnInit(): void {
        this.list$ = this.adminUsersService.getList().pipe(share());
    }
}
