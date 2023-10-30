import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { CompanyFormComponent } from '@app/companies/form/company-form.component';

import { CreateCompanyFormGroup } from './common/models/create-company.model';
import { CreateCompanyState, CreateCompanyStore } from './create-company.store';

@Component({
    selector: 'app-create-company',
    standalone: true,
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        FormErrorMessagePipe,
        RouterModule,
        LetDirective,
        CompanyFormComponent
    ],
    providers: [CreateCompanyStore],
    templateUrl: './create-company.component.html'
})
export class CreateCompanyComponent implements OnInit {
    form!: FormGroup<CreateCompanyFormGroup>;

    vm$: Observable<CreateCompanyState> = this.createCompanyStore.vm$;

    constructor(
        private readonly fb: FormBuilder,
        private readonly createCompanyStore: CreateCompanyStore
    ) {}

    ngOnInit(): void {
        this.initForm();
    }

    onSubmit(): void {
        const formData = new FormData();
        const formValue = this.form.getRawValue();
        for (const key of Object.keys(formValue)) {
            formData.append(key, formValue[key as keyof CreateCompanyFormGroup] as string);
        }
        this.createCompanyStore.submitForm(formData);
    }

    private initForm(): void {
        this.form = this.fb.group({
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
}
