import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';

import { provideComponentStore } from '@ngrx/component-store';
import { Store } from '@ngrx/store';

import { UsersListState, UsersListStore } from '@app/admin/users/list/users-list.store';
import { AuthActions } from '@app/common/stores/auth/auth.actions';
import { authFeature } from '@app/common/stores/auth/auth.reducer';

import { UserCardComponent } from './user-card/user-card.component';

@Component({
    selector: 'app-users-list',
    standalone: true,
    imports: [CommonModule, UserCardComponent],
    providers: [provideComponentStore(UsersListStore)],
    templateUrl: './users-list.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UsersListComponent implements OnInit {
    loggedIn$!: Observable<boolean>;

    readonly vm$: Observable<UsersListState> = this.usersListStore.vm$;

    constructor(
        private readonly store: Store,
        private readonly usersListStore: UsersListStore
    ) {}

    ngOnInit(): void {
        this.loadData();
        this.initSubscription();
    }

    onLogin(): void {
        this.store.dispatch(AuthActions.login());
    }

    private loadData(): void {
        this.usersListStore.getList();
    }

    private initSubscription(): void {
        this.loggedIn$ = this.store.select(authFeature.selectLoggedIn);
    }
}
