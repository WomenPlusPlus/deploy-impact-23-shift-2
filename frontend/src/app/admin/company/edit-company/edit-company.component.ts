import { Observable, take } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';

import { CompanyFormComponent } from '@app/admin/company/form/company-form.component';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';

import { CreateCompanyFormGroup } from '../create-company/common/models/create-company.model';
import { CreateCompanyState } from '../create-company/create-company.store';
import { EditCompanyStore } from './edit-company.store';

@Component({
    selector: 'app-edit-company',
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
    providers: [EditCompanyStore],
    templateUrl: './edit-company.component.html'
})
export class EditCompanyComponent implements OnInit {
    id?: number;
    form!: FormGroup<CreateCompanyFormGroup>;
    vm$: Observable<CreateCompanyState> = this.editCompanyStore.vm$;
    previousName = '';
    previousPhotoURL = 'assets/profile-picture-default-form.png';

    constructor(
        private readonly fb: FormBuilder,
        private readonly editCompanyStore: EditCompanyStore,
        private readonly route: ActivatedRoute,
        private readonly router: Router
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
        this.editCompanyStore.submitForm(formData);
    }

    private initForm(): void {
        this.form = this.fb.group({
            name: this.fb.control('', [Validators.required, Validators.maxLength(256)]),
            address: this.fb.control('', [Validators.required, Validators.maxLength(256)]),
            logo: this.fb.control(new File([], '')),
            linkedin: this.fb.control('', [Validators.required, Validators.maxLength(512)]),
            kununu: this.fb.control('', [Validators.required, Validators.maxLength(512)]),
            phone: this.fb.control('', [Validators.required, Validators.maxLength(256)]),
            email: this.fb.control('', [Validators.required, Validators.maxLength(512)]),
            mission: this.fb.control('', [Validators.required, Validators.maxLength(1024)]),
            values: this.fb.control('', [Validators.required, Validators.maxLength(1024)]),
            jobtypes: this.fb.control('', [Validators.required, Validators.maxLength(1024)]),
            expectation: this.fb.control('', [Validators.required, Validators.maxLength(1024)])
        });

        this.id = Number(this.route.snapshot.paramMap.get('id'));
        if (!this.id) {
            this.router.navigate(['/admin', 'companies']);
            return;
        }

        this.editCompanyStore
            .getValues(this.id)
            .pipe(take(1))
            .subscribe((data) => {
                this.form.patchValue({
                    name: data.name,
                    address: data.address,
                    logo: null,
                    linkedin: data.linkedin,
                    kununu: data.kununu,
                    phone: data.phone,
                    email: data.email,
                    mission: data.mission,
                    values: data.values,
                    jobtypes: data.jobtypes,
                    expectation: data.expectation
                });
                this.previousName = data.name;
                this.previousPhotoURL = data.logo;
            });
    }
}
