import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, ViewChild } from '@angular/core';
import { FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { CreateUserAssociationComponent } from '@app/admin/users/creation/association/create-user-association.component';
import { CreateUserCandidateComponent } from '@app/admin/users/creation/candidate/create-user-candidate.component';
import { CreateUserFormGroup } from '@app/admin/users/creation/common/models/create-user.model';
import { CreateUserCompanyComponent } from '@app/admin/users/creation/company/create-user-company.component';
import { CreateUserStore } from '@app/admin/users/creation/create-user.store';
import { CreateUserGenericComponent } from '@app/admin/users/creation/generic/create-user-generic.component';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { UserKindEnum } from '@app/common/models/users.model';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';

type CreateUserFormComponent = { form: FormGroup<CreateUserFormGroup> };

@Component({
    selector: 'app-create-user',
    standalone: true,
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        RouterModule,
        CreateUserGenericComponent,
        CreateUserCandidateComponent,
        FormErrorMessagePipe,
        LetDirective,
        UserKindLabelPipe,
        CreateUserCompanyComponent,
        CreateUserAssociationComponent
    ],
    providers: [CreateUserStore],
    templateUrl: './create-user.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateUserComponent {
    @ViewChild('childFormEl', { static: false })
    set childFormComponent(el: CreateUserFormComponent | undefined) {
        this.childForm = el?.form;
    }

    vm$ = this.createUserStore.vm$;
    childForm?: FormGroup<CreateUserFormGroup>;
    selectedKind = UserKindEnum.CANDIDATE;

    get userEmail(): string {
        return this.childForm?.controls.details.controls.email.value || 'N/A';
    }

    protected readonly userKindsEnum = UserKindEnum;
    protected readonly userKinds: UserKindEnum[] = [
        UserKindEnum.ADMIN,
        UserKindEnum.CANDIDATE,
        UserKindEnum.COMPANY,
        UserKindEnum.ASSOCIATION
    ];

    constructor(private readonly createUserStore: CreateUserStore) {}

    onSubmit(): void {
        if (!this.childForm?.valid) {
            return;
        }
        this.createUserStore.submitForm(this.childForm.getRawValue());
    }
}
