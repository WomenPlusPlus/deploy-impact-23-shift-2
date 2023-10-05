import { Observable, share } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';

import { Store } from '@ngrx/store';

import { AuthActions } from '@app/common/stores/auth/auth.actions';
import { authFeature } from '@app/common/stores/auth/auth.reducer';

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
    loggedIn$!: Observable<boolean>;

    constructor(
        private readonly store: Store,
        private readonly adminUsersService: AdminUsersService
    ) {}

    ngOnInit(): void {
        this.initSubscription();
    }

    onLogin(): void {
        this.store.dispatch(AuthActions.login());
    }

    private initSubscription(): void {
        this.list$ = this.adminUsersService.getList().pipe(share());
        this.loggedIn$ = this.store.select(authFeature.selectLoggedIn);
    }
}
