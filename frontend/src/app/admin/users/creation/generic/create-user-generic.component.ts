import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { map, Observable, startWith } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

import { CreateUserFormGroup } from '@app/admin/users/creation/common/models/create-user.model';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';

const DEFAULT_PHOTO_URL = 'assets/profile-picture-default-creation.png';

@Component({
    selector: 'app-create-user-generic',
    standalone: true,
    imports: [CommonModule, FontAwesomeModule, FormErrorMessagePipe, LetDirective, ReactiveFormsModule],
    templateUrl: './create-user-generic.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateUserGenericComponent implements OnInit {
    form!: FormGroup<CreateUserFormGroup>;
    imagePreviewUrl$!: Observable<string>;

    get detailsForm(): CreateUserFormGroup['details'] {
        return this.form.controls.details;
    }

    get socialForm(): CreateUserFormGroup['social'] {
        return this.form.controls.social;
    }

    constructor(private readonly fb: FormBuilder) {}

    ngOnInit(): void {
        this.initForm();
        this.initSubscriptions();
    }

    onPhotoUpload(event: Event): void {
        const file = (event.target as HTMLInputElement).files?.[0];
        this.detailsForm.controls.photo.setValue(file || null);
    }

    private initForm(): void {
        this.form = this.fb.group({
            details: this.fb.group({
                firstName: this.fb.control<string | null>(null, [
                    Validators.required,
                    Validators.minLength(3),
                    Validators.maxLength(128)
                ]),
                lastName: this.fb.control<string | null>(null, [
                    Validators.required,
                    Validators.minLength(3),
                    Validators.maxLength(128)
                ]),
                preferredName: this.fb.control<string | null>(null, [
                    Validators.minLength(3),
                    Validators.maxLength(256)
                ]),
                email: this.fb.control<string | null>(null, [
                    Validators.required,
                    Validators.minLength(5),
                    Validators.maxLength(512)
                ]),
                phoneNumber: this.fb.control<string | null>(null, [
                    Validators.required,
                    Validators.minLength(3),
                    Validators.maxLength(20)
                ]),
                birthDate: this.fb.control<Date | null>(null, [Validators.required]),
                photo: this.fb.control<File | null>(null)
            }),
            social: this.fb.group({
                linkedInUrl: this.fb.control<string | null>(null),
                githubUrl: this.fb.control<string | null>(null),
                portfolioUrl: this.fb.control<string | null>(null)
            })
        });
    }

    private initSubscriptions(): void {
        this.imagePreviewUrl$ = this.detailsForm.controls.photo.valueChanges.pipe(
            startWith(this.detailsForm.controls.photo.value),
            map((file: File | null) => {
                if (!file) {
                    return DEFAULT_PHOTO_URL;
                }
                try {
                    return URL.createObjectURL(file);
                } catch (error) {
                    console.error(error);
                    return DEFAULT_PHOTO_URL;
                }
            })
        );
    }
}
