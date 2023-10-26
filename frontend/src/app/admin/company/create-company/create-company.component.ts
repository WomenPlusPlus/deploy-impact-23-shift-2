import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { CompanyFormComponent } from '@app/admin/company/form/company-form.component';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';

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
            formData.append(key, formValue[key as keyof CreateCompanyFormGroup]!);
        }
        this.createCompanyStore.submitForm(formData);
    }

    onFileSelected(event: Event): void {
        const file = (event.target as HTMLInputElement).files?.[0];
        this.form.controls.logo.setValue(file || null);
    }

    private initForm(): void {
        this.form = this.fb.group({
            name: this.fb.control('', [Validators.required, Validators.maxLength(256)]),
            address: this.fb.control('', [Validators.required, Validators.maxLength(256)]),
            logo: this.fb.control(new File([], ''), Validators.required),
            linkedin: this.fb.control('', [Validators.required, Validators.maxLength(512)]),
            kununu: this.fb.control('', [Validators.required, Validators.maxLength(512)]),
            phone: this.fb.control('', [Validators.required, Validators.maxLength(256)]),
            email: this.fb.control('', [Validators.required, Validators.maxLength(512)]),
            mission: this.fb.control('', [Validators.required, Validators.maxLength(1024)]),
            values: this.fb.control('', [Validators.required, Validators.maxLength(1024)]),
            jobtypes: this.fb.control('', [Validators.required, Validators.maxLength(1024)]),
            expectation: this.fb.control('', [Validators.required, Validators.maxLength(1024)])
        });
    }
}
