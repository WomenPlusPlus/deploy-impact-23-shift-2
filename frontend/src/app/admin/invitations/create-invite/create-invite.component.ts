import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

import { CreateInviteFormModel } from '@app/admin/invitations/create-invite/common/models/create-invite.model';
import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';

@Component({
    selector: 'app-create-invite',
    standalone: true,
    imports: [CommonModule, FormsModule, ReactiveFormsModule, UserKindLabelPipe, UserCompanyRoleLabelPipe],
    templateUrl: './create-invite.component.html'
})
export class CreateInviteComponent implements OnInit {
    form!: CreateInviteFormModel;

    readonly userKinds: UserKindEnum[] = [
        UserKindEnum.ADMIN,
        UserKindEnum.CANDIDATE,
        UserKindEnum.COMPANY,
        UserKindEnum.ASSOCIATION
    ];
    readonly userRoles: UserRoleEnum[] = [UserRoleEnum.ADMIN, UserRoleEnum.USER];

    protected readonly userKindEnum = UserKindEnum;

    constructor(private router: Router) {}

    ngOnInit(): void {
        this.initForm();
    }

    onSubmit(): void {
        this.router.navigate(['/']);
    }

    private initForm(): void {
        this.form = {
            kind: UserKindEnum.CANDIDATE,
            role: UserRoleEnum.ADMIN,
            email: '',
            subject: '',
            message: ''
        };
    }
}
