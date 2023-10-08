import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faCog } from '@fortawesome/free-solid-svg-icons';
import { map, Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { UsersListStore } from '@app/admin/users/list/users-list.store';
import { UserKindEnum } from '@app/common/models/users.model';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';

@Component({
    selector: 'app-users-filters',
    standalone: true,
    imports: [CommonModule, FormsModule, FontAwesomeModule, UserKindLabelPipe],
    templateUrl: './users-filters.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UsersFiltersComponent implements OnInit {
    mode$ = this.usersListStore.mode$;
    searchTerm$ = this.usersListStore.filterSearchTerm$;

    protected readonly faCog = faCog;
    protected readonly userKindsOptions$: Observable<{ label: UserKindEnum; value: boolean }[]> =
        this.usersListStore.filterKinds$.pipe(
            map((selected: { [kind in UserKindEnum]: boolean }) =>
                this.userKinds.map((kind) => ({ label: kind, value: selected[kind] }))
            )
        );
    private readonly userKinds: UserKindEnum[] = [
        UserKindEnum.ADMIN,
        UserKindEnum.CANDIDATE,
        UserKindEnum.COMPANY,
        UserKindEnum.ASSOCIATION
    ];

    constructor(private readonly usersListStore: UsersListStore) {}

    ngOnInit(): void {
        this.initData();
    }

    onKindChange(kind: UserKindEnum): void {
        this.usersListStore.updateFilterKind(kind);
    }

    onSearchTermChange(term: string): void {
        this.usersListStore.updateFilterSearchTerm(term);
    }

    onDetailedModeChange(event: boolean): void {
        this.usersListStore.setMode(event ? 'detailed' : 'short');
    }

    private initData(): void {
        this.usersListStore.initFilters();
    }
}
