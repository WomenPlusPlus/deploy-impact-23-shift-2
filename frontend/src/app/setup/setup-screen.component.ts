import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';

import { SetupAssociationUserFormComponent } from '@app/setup/association/setup-association-user-form.component';
import { SetupCandidateFormComponent } from '@app/setup/candidate/setup-candidate-form.component';
import { SetupCompanyUserFormComponent } from '@app/setup/company/setup-company-user-form.component';

@Component({
    selector: 'app-setup-screen',
    standalone: true,
    imports: [
        CommonModule,
        SetupCandidateFormComponent,
        SetupCompanyUserFormComponent,
        SetupAssociationUserFormComponent
    ],
    templateUrl: './setup-screen.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SetupScreenComponent {}
