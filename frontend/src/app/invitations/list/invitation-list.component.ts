import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { InvitationListState, InvitationListStore } from '@app/admin/invitations/list/invitation-list.store';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { InviteStatusLabelPipe } from '@app/common/pipes/invite-status-label/invite-status-label.pipe';
import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';
import { UserBadgesComponent } from '@app/ui/user-badges/user-badges.component';

@Component({
    selector: 'app-invitation-list',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        ContentErrorComponent,
        ContentLoadingComponent,
        FormsModule,
        JobLocationTypePipe,
        JobTypePipe,
        LetDirective,
        UserBadgesComponent,
        InviteStatusLabelPipe
    ],
    providers: [InvitationListStore],
    templateUrl: './invitation-list.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class InvitationListComponent implements OnInit {
    vm$: Observable<InvitationListState> = this.invitationListStore.vm$;

    constructor(private readonly invitationListStore: InvitationListStore) {}

    ngOnInit(): void {
        this.loadData();
    }

    onMyInvitesChange(state: boolean): void {
        this.invitationListStore.updateFilterOnlyMine(state);
    }

    onSearchTermChange(term: string): void {
        this.invitationListStore.updateFilterSearchTerm(term);
    }

    private loadData(): void {
        this.invitationListStore.initFilters();
        this.invitationListStore.getList();
    }
}
