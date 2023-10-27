import { merge, Observable, startWith, Subscription } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnDestroy, OnInit } from '@angular/core';
import {
    AbstractControl,
    FormBuilder,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    ValidatorFn,
    Validators
} from '@angular/forms';
import { RouterModule } from '@angular/router';

import { CreateInviteFormGroup } from '@app/admin/invitations/create-invite/common/models/create-invite.model';
import { CreateNewOptionPipe } from '@app/admin/invitations/create-invite/common/pipes/create-new-option/create-new-option.pipe';
import { CreateInviteState, CreateInviteStore } from '@app/admin/invitations/create-invite/create-invite.store';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { SelectSingleComponent } from '@app/ui/select-single/select-single.component';

const DEFAULT_INVITATION_SUBJECT = 'You have been invited to Shift2';
const DEFAULT_INVITATION_MESSAGE = `We are very excited for you to join our platform Shift2 - a gateway to discover new professional opportunities.\n
Click the link below to register your account and you can start exploring, connecting and expanding your network.\n
Thanks,
Shift2 Team`;

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
        LetDirective,
        SelectSingleComponent,
        CreateNewOptionPipe
    ],
    providers: [CreateInviteStore],
    templateUrl: './create-invite.component.html'
})
export class CreateInviteComponent implements OnInit, OnDestroy {
    form!: FormGroup<CreateInviteFormGroup>;

    vm$: Observable<CreateInviteState> = this.createInviteStore.vm$;

    protected readonly userKinds: UserKindEnum[] = [
        UserKindEnum.ADMIN,
        UserKindEnum.CANDIDATE,
        UserKindEnum.COMPANY,
        UserKindEnum.ASSOCIATION
    ];
    protected userKindEnum = UserKindEnum;
    protected readonly userRoles: UserRoleEnum[] = [UserRoleEnum.ADMIN, UserRoleEnum.USER];
    protected userRoleEnum = UserRoleEnum;

    private readonly subscriptions: Subscription[] = [];

    constructor(
        private readonly fb: FormBuilder,
        private readonly createInviteStore: CreateInviteStore
    ) {}

    ngOnInit(): void {
        this.loadData();
        this.initForm();
        this.initSubscriptions();
    }

    ngOnDestroy(): void {
        this.subscriptions.forEach((s) => s.unsubscribe());
    }

    onSubmit(): void {
        this.createInviteStore.submitForm(this.form.getRawValue());
    }

    private loadData(): void {
        this.createInviteStore.loadData();
    }

    private initForm(): void {
        this.form = this.fb.group({
            kind: this.fb.control(UserKindEnum.CANDIDATE, [Validators.required]),
            role: this.fb.control(UserRoleEnum.ADMIN, [Validators.required]),
            companyId: this.fb.control<number | null>(null, [
                this.requiredIf(
                    () =>
                        this.form?.controls.kind.value === UserKindEnum.COMPANY &&
                        this.form.controls.role.value !== UserRoleEnum.ADMIN
                )
            ]),
            associationId: this.fb.control<number | null>(null, [
                this.requiredIf(
                    () =>
                        this.form?.controls.kind.value === UserKindEnum.ASSOCIATION &&
                        this.form.controls.role.value !== UserRoleEnum.ADMIN
                )
            ]),
            email: this.fb.control('', [Validators.required, Validators.minLength(5), Validators.maxLength(512)]),
            subject: this.fb.control(DEFAULT_INVITATION_SUBJECT, [
                Validators.required,
                Validators.minLength(3),
                Validators.maxLength(256)
            ]),
            message: this.fb.control(DEFAULT_INVITATION_MESSAGE, [
                Validators.required,
                Validators.minLength(25),
                Validators.maxLength(1024)
            ])
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
        this.subscriptions.push(
            merge(this.form.controls.kind.valueChanges, this.form.controls.role.valueChanges).subscribe(() => {
                this.form.controls.associationId.reset();
                this.form.controls.companyId.reset();
            })
        );
    }

    private requiredIf(conditionFn: (control: AbstractControl) => boolean): ValidatorFn {
        return (control: AbstractControl) => (conditionFn(control) ? Validators.required(control) : null);
    }
}
