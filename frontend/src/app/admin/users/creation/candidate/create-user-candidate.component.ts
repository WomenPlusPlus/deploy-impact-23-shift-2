import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faAdd, faRemove } from '@fortawesome/free-solid-svg-icons';
import { map, Observable, startWith } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

import {
    CandidateEducationHistoryFormGroup,
    CandidateEducationHistoryFormModel,
    CandidateEmploymentHistoryFormGroup,
    CandidateEmploymentHistoryFormModel,
    CandidateSkillsFormGroup,
    CandidateSkillsFormModel,
    CandidateSpokenLanguagesFormGroup,
    CandidateSpokenLanguagesFormModel,
    CreateUserCandidateFormGroup
} from '@app/admin/users/creation/common/models/create-user.model';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';

const DEFAULT_PHOTO_URL = 'assets/profile-picture-default-creation.png';

@Component({
    selector: 'app-create-user-candidate',
    standalone: true,
    imports: [
        CommonModule,
        ReactiveFormsModule,
        FormErrorMessagePipe,
        LetDirective,
        ReactiveFormsModule,
        UserCompanyRoleLabelPipe,
        UserKindLabelPipe,
        FontAwesomeModule
    ],
    templateUrl: './create-user-candidate.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateUserCandidateComponent implements OnInit {
    form!: FormGroup<CreateUserCandidateFormGroup>;
    spokenLanguagesForm!: FormGroup<CandidateSpokenLanguagesFormGroup>;
    skillsForm!: FormGroup<CandidateSkillsFormGroup>;
    educationHistoryForm!: FormGroup<CandidateEducationHistoryFormGroup>;
    employmentHistoryForm!: FormGroup<CandidateEmploymentHistoryFormGroup>;
    imagePreviewUrl$!: Observable<string>;

    get detailsForm(): CreateUserCandidateFormGroup['details'] {
        return this.form.controls.details;
    }

    get jobForm(): CreateUserCandidateFormGroup['job'] {
        return this.form.controls.job;
    }

    get technicalForm(): CreateUserCandidateFormGroup['technical'] {
        return this.form.controls.technical;
    }

    get socialForm(): CreateUserCandidateFormGroup['social'] {
        return this.form.controls.social;
    }

    get spokenLanguages(): CandidateSpokenLanguagesFormModel[] {
        return this.technicalForm.controls.spokenLanguages.value || [];
    }

    get skills(): CandidateSkillsFormModel[] {
        return this.technicalForm.controls.skills.value || [];
    }

    get educationHistory(): CandidateEducationHistoryFormModel[] {
        return this.technicalForm.controls.educationHistory.value || [];
    }

    get employmentHistory(): CandidateEmploymentHistoryFormModel[] {
        return this.technicalForm.controls.employmentHistory.value || [];
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

    onCVUpload(event: Event): void {
        const file = (event.target as HTMLInputElement).files?.[0];
        this.technicalForm.controls.cv.setValue(file || null);
    }

    onAttachmentsUpload(event: Event): void {
        const files = (event.target as HTMLInputElement).files;
        this.technicalForm.controls.attachments.setValue(Array.from(files || []));
    }

    onVideoUpload(event: Event): void {
        const file = (event.target as HTMLInputElement).files?.[0];
        this.technicalForm.controls.video.setValue(file || null);
    }

    onAddSpokenLanguage(): void {
        const control = this.spokenLanguagesForm;
        if (control.invalid) {
            return;
        }
        this.technicalForm.controls.spokenLanguages.patchValue([...this.spokenLanguages, control.getRawValue()]);
        control.markAsPristine();
        control.reset();
    }

    onRemoveSpokenLanguage(index: number): void {
        const languages = this.spokenLanguages;
        this.technicalForm.controls.spokenLanguages.patchValue(
            languages.slice(0, index).concat(languages.slice(index + 1))
        );
    }

    onAddSkill(): void {
        const control = this.skillsForm;
        if (control.invalid) {
            return;
        }
        this.technicalForm.controls.skills.patchValue([...this.skills, control.getRawValue()]);
        control.markAsPristine();
        control.reset();
    }

    onRemoveSkill(index: number): void {
        const skills = this.skills;
        this.technicalForm.controls.skills.patchValue(skills.slice(0, index).concat(skills.slice(index + 1)));
    }

    onGoingEducationCheck(event: Event): void {
        const checked = (event.target as HTMLInputElement).checked;
        const control = this.educationHistoryForm.controls.toDate;
        if (!checked) {
            control.enable();
            return;
        }
        control.setValue(null);
        control.disable();
    }

    onAddEducation(): void {
        const control = this.educationHistoryForm;
        if (control.invalid) {
            return;
        }
        this.technicalForm.controls.educationHistory.patchValue([...this.educationHistory, control.getRawValue()]);
        control.controls.toDate.enable();
        control.markAsPristine();
        control.reset();
    }

    onRemoveEducation(index: number): void {
        const educationHistory = this.educationHistory;
        this.technicalForm.controls.educationHistory.patchValue(
            educationHistory.slice(0, index).concat(educationHistory.slice(index + 1))
        );
    }

    onGoingEmploymentCheck(event: Event): void {
        const checked = (event.target as HTMLInputElement).checked;
        const control = this.employmentHistoryForm.controls.toDate;
        if (!checked) {
            control.enable();
            return;
        }
        control.setValue(null);
        control.disable();
    }

    onAddEmployment(): void {
        const control = this.employmentHistoryForm;
        if (control.invalid) {
            return;
        }
        this.technicalForm.controls.employmentHistory.patchValue([...this.employmentHistory, control.getRawValue()]);
        control.controls.toDate.enable();
        control.markAsPristine();
        control.reset();
    }

    onRemoveEmployment(index: number): void {
        const employmentHistory = this.employmentHistory;
        this.technicalForm.controls.employmentHistory.patchValue(
            employmentHistory.slice(0, index).concat(employmentHistory.slice(index + 1))
        );
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
                    Validators.required,
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
            job: this.fb.group({
                yearsOfExperience: this.fb.control<number | null>(null, [Validators.required, Validators.min(0)]),
                jobStatus: this.fb.control<string | null>(null, [Validators.required]),
                seekJobType: this.fb.control<string | null>(null),
                seekCompanySize: this.fb.control<string | null>(null),
                seekLocations: this.fb.control<string[] | null>(null, [Validators.required, Validators.max(5)]),
                seekLocationType: this.fb.control<string | null>(null, [Validators.required]),
                seekSalary: this.fb.control<number | null>(null),
                seekValues: this.fb.control<string | null>(null),
                workPermit: this.fb.control<string | null>(null, [Validators.required]),
                noticePeriod: this.fb.control<string | null>(null)
            }),
            technical: this.fb.group({
                spokenLanguages: this.fb.control<CandidateSpokenLanguagesFormModel[] | null>([]),
                skills: this.fb.control<CandidateSkillsFormModel[] | null>([]),
                cv: this.fb.control<File | null>(null),
                attachments: this.fb.control<File[] | null>([]),
                video: this.fb.control<File | null>(null),
                educationHistory: this.fb.control<CandidateEducationHistoryFormModel[] | null>([]),
                employmentHistory: this.fb.control<CandidateEmploymentHistoryFormModel[] | null>([])
            }),
            social: this.fb.group({
                linkedInUrl: this.fb.control<string | null>(null),
                githubUrl: this.fb.control<string | null>(null),
                portfolioUrl: this.fb.control<string | null>(null)
            })
        });
        this.spokenLanguagesForm = this.fb.group({
            language: this.fb.control<string | null>(null, [Validators.required]),
            level: this.fb.control<number | null>(null, [Validators.required, Validators.min(0), Validators.max(5)])
        });
        this.skillsForm = this.fb.group({
            name: this.fb.control<string | null>(null, [Validators.required]),
            years: this.fb.control<number | null>(null, [Validators.required, Validators.min(0)])
        });
        this.educationHistoryForm = this.fb.group({
            title: this.fb.control<string | null>(null, [
                Validators.required,
                Validators.minLength(3),
                Validators.maxLength(128)
            ]),
            description: this.fb.control<string | null>(null, [
                Validators.required,
                Validators.minLength(3),
                Validators.maxLength(512)
            ]),
            entity: this.fb.control<string | null>(null, [
                Validators.required,
                Validators.minLength(3),
                Validators.maxLength(128)
            ]),
            fromDate: this.fb.control<Date | null>(null, [Validators.required]),
            toDate: this.fb.control<Date | null>(null),
            onGoing: this.fb.control<boolean | null>(false)
        });
        this.employmentHistoryForm = this.fb.group({
            title: this.fb.control<string | null>(null, [
                Validators.required,
                Validators.minLength(3),
                Validators.maxLength(128)
            ]),
            description: this.fb.control<string | null>(null, [
                Validators.required,
                Validators.minLength(3),
                Validators.maxLength(512)
            ]),
            company: this.fb.control<string | null>(null, [
                Validators.required,
                Validators.minLength(3),
                Validators.maxLength(128)
            ]),
            fromDate: this.fb.control<Date | null>(null, [Validators.required]),
            toDate: this.fb.control<Date | null>(null),
            onGoing: this.fb.control<boolean | null>(false)
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

    protected readonly faAdd = faAdd;
    protected readonly faRemove = faRemove;
}
