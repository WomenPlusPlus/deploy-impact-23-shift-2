import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';

import { UserKindEnum } from '@app/common/models/users.model';
import { SetupAdminFormComponent } from '@app/setup/admin/setup-admin-form.component';
import { SetupAssociationUserFormComponent } from '@app/setup/association/setup-association-user-form.component';
import { SetupCandidateFormComponent } from '@app/setup/candidate/setup-candidate-form.component';
import { SetupCompanyUserFormComponent } from '@app/setup/company/setup-company-user-form.component';
import { SetupScreenStore } from '@app/setup/setup-screen.store';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

@Component({
    selector: 'app-setup-screen',
    standalone: true,
    imports: [
        CommonModule,
        SetupCandidateFormComponent,
        SetupCompanyUserFormComponent,
        SetupAssociationUserFormComponent,
        SetupAdminFormComponent,
        ContentErrorComponent,
        ContentLoadingComponent
    ],
    providers: [SetupScreenStore],
    templateUrl: './setup-screen.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SetupScreenComponent implements OnInit {
    vm$ = this.setupScreenStore.vm$;

    protected readonly userKindEnum = UserKindEnum;

    constructor(private readonly setupScreenStore: SetupScreenStore) {}

    ngOnInit(): void {
        this.loadData();
    }

    private loadData(): void {
        this.setupScreenStore.getProfile();
    }
}
