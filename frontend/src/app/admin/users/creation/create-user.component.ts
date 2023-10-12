import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, ViewChild } from '@angular/core';
import { FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';

import { CreateUserCandidateComponent } from '@app/admin/users/creation/candidate/create-user-candidate.component';
import { CreateUserFormGroup } from '@app/admin/users/creation/common/models/create-user.model';
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
        CreateUserGenericComponent,
        CreateUserCandidateComponent,
        FormErrorMessagePipe,
        LetDirective,
        ReactiveFormsModule,
        UserKindLabelPipe
    ],
    templateUrl: './create-user.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateUserComponent {
    @ViewChild('childFormEl', { static: false })
    set childFormComponent(el: CreateUserFormComponent | undefined) {
        this.childForm = el?.form;
    }

    selectedKind = UserKindEnum.CANDIDATE;

    childForm?: FormGroup<CreateUserFormGroup>;

    protected readonly userKindsEnum = UserKindEnum;
    protected readonly userKinds: UserKindEnum[] = [
        UserKindEnum.ADMIN,
        UserKindEnum.CANDIDATE,
        UserKindEnum.COMPANY,
        UserKindEnum.ASSOCIATION
    ];

    onSubmit(): void {
        if (!this.childForm?.valid) {
            return;
        }
    }
}
