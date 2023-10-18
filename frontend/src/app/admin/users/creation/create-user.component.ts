import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, ViewChild } from '@angular/core';
import { FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { CreateUserAssociationComponent } from '@app/admin/users/creation/association/create-user-association.component';
import { CreateUserCandidateComponent } from '@app/admin/users/creation/candidate/create-user-candidate.component';
import { CreateUserFormGroup, CreateUserFormModel } from '@app/admin/users/creation/common/models/create-user.model';
import { CreateUserCompanyComponent } from '@app/admin/users/creation/company/create-user-company.component';
import { CreateUserStore } from '@app/admin/users/creation/create-user.store';
import { CreateUserGenericComponent } from '@app/admin/users/creation/generic/create-user-generic.component';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { UserKindEnum } from '@app/common/models/users.model';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';

export interface CreateUserFormComponent<T extends CreateUserFormGroup = any, S extends CreateUserFormModel = any> {
    form: FormGroup<T>;
    formValue: S;
}

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
    childFormComponent?: CreateUserFormComponent<CreateUserFormGroup, CreateUserFormModel>;

    vm$ = this.createUserStore.vm$;
    selectedKind = UserKindEnum.CANDIDATE;

    get childForm(): FormGroup<CreateUserFormGroup> | undefined {
        return this.childFormComponent?.form;
    }

    get userEmail(): string {
        return this.childFormComponent?.form.controls.details.controls.email.value || 'N/A';
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
        if (!this.childFormComponent?.form.valid) {
            return;
        }
        this.createUserStore.submitForm({
            ...this.childFormComponent.formValue,
            kind: this.selectedKind
        });
    }
}
