import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';

import { CreateAssociationFormGroup } from './common/models/create-association.model';
import { CreateAssociationState, CreateAssociationStore } from './create-association.store';

@Component({
    selector: 'app-create-association',
    standalone: true,
    imports: [CommonModule, FormsModule, ReactiveFormsModule, FormErrorMessagePipe, RouterModule, LetDirective],
    providers: [CreateAssociationStore],
    templateUrl: './create-association.component.html'
})
export class CreateAssociationComponent implements OnInit {
    form!: FormGroup<CreateAssociationFormGroup>;

    vm$: Observable<CreateAssociationState> = this.createAssociationStore.vm$;

    selectedFile: File | null = null;

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

    onFileSelected(event: Event): void {
        const file = (event.target as HTMLInputElement).files?.[0];
        this.form.controls.logo.setValue(file || null);
    }

    private initForm(): void {
        this.form = this.fb.group({
            name: this.fb.control('', [Validators.required, Validators.maxLength(256)]),
            logo: this.fb.control(new File([], ''), Validators.required),
            url: this.fb.control('', [Validators.required, Validators.maxLength(512)]),
            focus: this.fb.control('', [Validators.required, Validators.maxLength(1024)])
        });
    }
}
