import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

import { CreateInviteFormModel } from '@app/admin/invitations/create-invite/common/models/create-invite.model';
import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';

@Component({
    selector: 'app-create-invite',
    standalone: true,
    imports: [CommonModule, FormsModule, ReactiveFormsModule],
    templateUrl: './create-invite.component.html'
})
export class CreateInviteComponent implements OnInit {
    form!: CreateInviteFormModel;

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
            role: UserRoleEnum.USER,
            email: '',
            subject: '',
            message: ''
        };
    }
}
