import { Observable, take } from 'rxjs';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';

import { AssociationFormComponent } from '@app/associations/form/association-form.component';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';

import { CreateAssociationFormGroup } from '../create-association/common/models/create-association.model';
import { CreateAssociationState } from '../create-association/create-association.store';
import { EditAssociationStore } from './edit-association.store';

@Component({
    selector: 'app-edit-association',
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
    providers: [EditAssociationStore],
    templateUrl: './edit-association.component.html'
})
export class EditAssociationComponent implements OnInit {
    id?: number;
    form!: FormGroup<CreateAssociationFormGroup>;
    vm$: Observable<CreateAssociationState> = this.editAssociationStore.vm$;
    previousName = '';
    previousPhotoURL?: string;

    constructor(
        private readonly fb: FormBuilder,
        private readonly editAssociationStore: EditAssociationStore,
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
            formData.append(key, formValue[key as keyof CreateAssociationFormGroup] as string);
        }
        this.editAssociationStore.submitForm(formData);
    }

    private initForm(): void {
        this.form = this.fb.group({
            name: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(256)]),
            logo: this.fb.control<File | null>(null),
            websiteUrl: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(512)]),
            focus: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(1024)])
        });

        this.id = Number(this.route.snapshot.paramMap.get('id'));
        if (!this.id) {
            this.router.navigate(['/associations']);
            return;
        }

        this.editAssociationStore
            .getValues(this.id)
            .pipe(take(1))
            .subscribe((data) => {
                this.form.patchValue({
                    name: data.name,
                    logo: null,
                    websiteUrl: data.websiteUrl,
                    focus: data.focus
                });
                this.previousName = data.name;
                this.previousPhotoURL = data.imageUrl?.url;
            });
    }
}
