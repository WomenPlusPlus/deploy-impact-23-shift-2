import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faEdit, faMessage, faRemove } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';

import { UserKindEnum, UserStateEnum } from '@app/common/models/users.model';
import { CompanySizePipe } from '@app/common/pipes/company-size/company-size.pipe';
import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobStatusPipe } from '@app/common/pipes/job-status/job-status.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { UserStateLabelPipe } from '@app/common/pipes/user-state-label/user-state-label.pipe';
import { WorkPermitPipe } from '@app/common/pipes/work-permit/work-permit.pipe';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';
import { ViewCandidateSectionsComponent } from '@app/users/view/candidate/view-candidate-sections.component';
import { ViewUserStore } from '@app/users/view/view-user.store';

@Component({
    selector: 'app-view-user',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        FontAwesomeModule,
        JobStatusPipe,
        WorkPermitPipe,
        JobTypePipe,
        CompanySizePipe,
        JobLocationTypePipe,
        UserStateLabelPipe,
        ContentErrorComponent,
        ContentLoadingComponent,
        ViewCandidateSectionsComponent
    ],
    providers: [ViewUserStore],
    templateUrl: './view-user.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ViewUserComponent implements OnInit {
    vm$ = this.viewUserStore.vm$;

    protected readonly faMessage = faMessage;
    protected readonly faRemove = faRemove;
    protected readonly faEdit = faEdit;
    protected readonly userKindEnum = UserKindEnum;
    protected readonly userStateEnum = UserStateEnum;

    constructor(
        private readonly router: Router,
        private readonly route: ActivatedRoute,
        private readonly viewUserStore: ViewUserStore
    ) {}

    ngOnInit(): void {
        this.loadData();
    }

    onDelete(id: number): void {
        this.viewUserStore.deleteItem(id);
    }

    private loadData(): void {
        const id = this.route.snapshot.paramMap.get('id');
        if (!id || isNaN(+id)) {
            this.router.navigate(['..']);
            return;
        }
        this.viewUserStore.getUser(+id);
    }
}
