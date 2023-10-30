import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faEye } from '@fortawesome/free-solid-svg-icons';
import { map, Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { UserKindEnum } from '@app/common/models/users.model';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { UsersListStore } from '@app/users/list/users-list.store';

@Component({
    selector: 'app-users-filters',
    standalone: true,
    imports: [CommonModule, FormsModule, FontAwesomeModule, UserKindLabelPipe, LetDirective],
    templateUrl: './users-filters.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UsersFiltersComponent implements OnInit {
    searchTerm$ = this.usersListStore.filterSearchTerm$;

    protected readonly faEye = faEye;
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

    onDetailedModeChange(): void {
        this.usersListStore.toggleMode();
    }

    private initData(): void {
        this.usersListStore.initFilters();
    }
}
