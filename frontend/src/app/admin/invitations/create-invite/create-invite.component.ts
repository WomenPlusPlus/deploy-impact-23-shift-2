import { Observable, startWith, Subscription } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { CreateInviteFormGroup } from '@app/admin/invitations/create-invite/common/models/create-invite.model';
import { CreateInviteState, CreateInviteStore } from '@app/admin/invitations/create-invite/create-invite.store';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';

@Component({
    selector: 'app-create-invite',
    standalone: true,
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        RouterModule,
        FormErrorMessagePipe,
        UserKindLabelPipe,
        UserCompanyRoleLabelPipe,
        LetDirective
    ],
    providers: [CreateInviteStore],
    templateUrl: './create-invite.component.html'
})
export class CreateInviteComponent implements OnInit, OnDestroy {
    form!: FormGroup<CreateInviteFormGroup>;

    vm$: Observable<CreateInviteState> = this.createInviteStore.vm$;

    readonly userKinds: UserKindEnum[] = [
        UserKindEnum.ADMIN,
        UserKindEnum.CANDIDATE,
        UserKindEnum.COMPANY,
        UserKindEnum.ASSOCIATION
    ];
    readonly userRoles: UserRoleEnum[] = [UserRoleEnum.ADMIN, UserRoleEnum.USER];

    protected readonly subscriptions: Subscription[] = [];

    constructor(
        private readonly fb: FormBuilder,
        private readonly createInviteStore: CreateInviteStore
    ) {}

    ngOnInit(): void {
        this.initForm();
        this.initSubscriptions();
    }

    ngOnDestroy(): void {
        this.subscriptions.forEach((s) => s.unsubscribe());
    }

    onSubmit(): void {
        this.createInviteStore.submitForm(this.form.getRawValue());
    }

    private initForm(): void {
        this.form = this.fb.group({
            kind: this.fb.control(UserKindEnum.CANDIDATE, [Validators.required]),
            role: this.fb.control(UserRoleEnum.ADMIN, [Validators.required]),
            email: this.fb.control('', [Validators.required, Validators.minLength(5), Validators.maxLength(512)]),
            subject: this.fb.control('', [Validators.required, Validators.minLength(3), Validators.maxLength(256)]),
            message: this.fb.control('', [Validators.required, Validators.minLength(25), Validators.maxLength(1024)])
        });
    }

    private initSubscriptions(): void {
        this.subscriptions.push(
            this.form.controls.kind.valueChanges
                .pipe(startWith(this.form.controls.kind.value))
                .subscribe((kind) =>
                    kind !== UserKindEnum.ASSOCIATION && kind !== UserKindEnum.COMPANY
                        ? this.form.controls.role.disable()
                        : this.form.controls.role.enable()
                )
        );
    }
}
