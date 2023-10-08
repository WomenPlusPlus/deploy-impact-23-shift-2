import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faAdd } from '@fortawesome/free-solid-svg-icons';
import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { RouterModule } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';
import { Store } from '@ngrx/store';

import { UsersFiltersComponent } from '@app/admin/users/list/users-filters/users-filters.component';
import { UsersListState, UsersListStore } from '@app/admin/users/list/users-list.store';
import { UserAssociationRoleLabelPipe } from '@app/common/pipes/user-association-role-label/user-association-role-label.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { authFeature } from '@app/common/stores/auth/auth.reducer';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { UserCardComponent } from './user-card/user-card.component';

@Component({
    selector: 'app-users-list',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        UserCardComponent,
        UsersFiltersComponent,
        UserAssociationRoleLabelPipe,
        UserCompanyRoleLabelPipe,
        UserKindLabelPipe,
        FontAwesomeModule,
        ContentErrorComponent,
        ContentLoadingComponent
    ],
    providers: [provideComponentStore(UsersListStore)],
    templateUrl: './users-list.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UsersListComponent implements OnInit {
    loggedIn$!: Observable<boolean>;

    readonly vm$: Observable<UsersListState> = this.usersListStore.vm$;

    protected readonly faAdd = faAdd;

    constructor(
        private readonly store: Store,
        private readonly usersListStore: UsersListStore
    ) {}

    ngOnInit(): void {
        this.loadData();
        this.initSubscription();
    }

    private loadData(): void {
        this.usersListStore.getList();
    }

    private initSubscription(): void {
        this.loggedIn$ = this.store.select(authFeature.selectLoggedIn);
    }
}
