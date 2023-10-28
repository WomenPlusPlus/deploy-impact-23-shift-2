import { CommonModule } from '@angular/common';
import { AfterViewInit, ChangeDetectionStrategy, Component, Input, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, Validators } from '@angular/forms';

import { CreateCompanyFormGroup } from '@app/admin/company/create-company/common/models/create-company.model';
import { CompanyFormComponent } from '@app/admin/company/form/company-form.component';
import {
    UserFormCompanyFormGroup,
    UserFormComponent,
    UserFormGroup,
    UserFormModel
} from '@app/admin/users/form/common/models/user-form.model';
import { UserFormGenericComponent } from '@app/admin/users/form/generic/user-form-generic.component';
import { UserFormStore } from '@app/admin/users/form/user-form.store';
import { ProfileSetup } from '@app/common/models/profile.model';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { SetupCompanyUserFormStore } from '@app/setup/company/setup-company-user-form.store';

@Component({
    selector: 'app-setup-company-user-form',
    standalone: true,
    imports: [CommonModule, UserKindLabelPipe, FormsModule, CompanyFormComponent, UserFormGenericComponent],
    providers: [SetupCompanyUserFormStore, UserFormStore],
    templateUrl: './setup-company-user-form.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SetupCompanyUserFormComponent implements OnInit, AfterViewInit {
    @Input() profile!: ProfileSetup;

    @ViewChild('childFormEl', { static: false })
    childFormComponent?: UserFormComponent<UserFormGroup, UserFormModel>;

    vm$ = this.setupCompanyUserStore.vm$;
    companyForm!: FormGroup<CreateCompanyFormGroup>;

    get childForm(): FormGroup<UserFormGroup> | undefined {
        return this.childFormComponent?.form;
    }

    constructor(
        private readonly fb: FormBuilder,
        private readonly setupCompanyUserStore: SetupCompanyUserFormStore
    ) {}

    ngOnInit(): void {
        this.initForm();
    }

    ngAfterViewInit(): void {
        this.initFormValue();
    }

    onSubmit(): void {
        if (!this.childFormComponent?.form.valid || !this.companyForm.valid) {
            return;
        }
        const companyData = new FormData();
        const formValue = this.companyForm.getRawValue();
        for (const key of Object.keys(formValue)) {
            companyData.append(key, formValue[key as keyof CreateCompanyFormGroup] as string);
        }
        this.setupCompanyUserStore.submitForm({
            companyId: this.profile.company?.id,
            user: {
                ...this.childFormComponent.formValue,
                kind: this.profile.kind,
                role: this.profile.role
            },
            company: companyData
        });
    }

    private initForm(): void {
        this.companyForm = this.fb.group({
            name: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(256)]),
            address: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(256)]),
            logo: this.fb.control<File | null>(null),
            linkedin: this.fb.control<string | null>(null, [Validators.maxLength(512)]),
            kununu: this.fb.control<string | null>(null, [Validators.maxLength(512)]),
            phone: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(256)]),
            email: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(512)]),
            mission: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(1024)]),
            values: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(1024)]),
            jobtypes: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(1024)]),
            expectation: this.fb.control<string | null>(null, [Validators.maxLength(1024)])
        });
    }

    private initFormValue(): void {
        const form = this.childForm as FormGroup<UserFormCompanyFormGroup> | undefined;
        if (!form || !this.profile) {
            return;
        }
        form.patchValue({ details: { email: this.profile.email } });
        form.controls.details.controls.email.disable({ emitEvent: false });

        if (!this.profile.company) {
            return;
        }
        const { logo, ...company } = this.profile.company;
        this.companyForm.patchValue(company);
        this.companyForm.disable({ emitEvent: false });
    }
}
