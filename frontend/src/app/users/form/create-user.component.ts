import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, ViewChild } from '@angular/core';
import { FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { UserFormAssociationComponent } from '@app/admin/users/form/association/user-form-association.component';
import { UserFormCandidateComponent } from '@app/admin/users/form/candidate/user-form-candidate.component';
import { UserFormComponent, UserFormGroup, UserFormModel } from '@app/admin/users/form/common/models/user-form.model';
import { UserFormCompanyComponent } from '@app/admin/users/form/company/user-form-company.component';
import { CreateUserStore } from '@app/admin/users/form/create-user.store';
import { UserFormGenericComponent } from '@app/admin/users/form/generic/user-form-generic.component';
import { UserFormStore } from '@app/admin/users/form/user-form.store';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { UserKindEnum } from '@app/common/models/users.model';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';

@Component({
    selector: 'app-create-user',
    standalone: true,
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        RouterModule,
        UserFormGenericComponent,
        UserFormCandidateComponent,
        FormErrorMessagePipe,
        LetDirective,
        UserKindLabelPipe,
        UserFormCompanyComponent,
        UserFormAssociationComponent
    ],
    providers: [UserFormStore, CreateUserStore],
    templateUrl: './create-user.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateUserComponent {
    @ViewChild('childFormEl', { static: false })
    childFormComponent?: UserFormComponent<UserFormGroup, UserFormModel>;

    vm$ = this.createUserStore.vm$;
    selectedKind = UserKindEnum.CANDIDATE;

    get childForm(): FormGroup<UserFormGroup> | undefined {
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
