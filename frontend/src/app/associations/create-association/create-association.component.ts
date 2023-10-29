import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { AssociationFormComponent } from '@app/associations/form/association-form.component';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';

import { CreateAssociationFormGroup } from './common/models/create-association.model';
import { CreateAssociationState, CreateAssociationStore } from './create-association.store';

@Component({
    selector: 'app-create-association',
    standalone: true,
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        FormErrorMessagePipe,
        RouterModule,
        LetDirective,
        AssociationFormComponent
    ],
    providers: [CreateAssociationStore],
    templateUrl: './create-association.component.html'
})
export class CreateAssociationComponent implements OnInit {
    form!: FormGroup<CreateAssociationFormGroup>;

    vm$: Observable<CreateAssociationState> = this.createAssociationStore.vm$;

    constructor(
        private readonly fb: FormBuilder,
        private readonly createAssociationStore: CreateAssociationStore
    ) {}

    ngOnInit(): void {
        this.initForm();
    }

    onSubmit(): void {
        const formData = new FormData();
        const formValue = this.form.getRawValue();
        for (const key of Object.keys(formValue)) {
            formData.append(key, formValue[key as keyof CreateAssociationFormGroup]!);
        }
        this.createAssociationStore.submitForm(formData);
    }

    private initForm(): void {
        this.form = this.fb.group({
            name: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(256)]),
            logo: this.fb.control<File | null>(null),
            websiteUrl: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(512)]),
            focus: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(1024)])
        });
    }
}
