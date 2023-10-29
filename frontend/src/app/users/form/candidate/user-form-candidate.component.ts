import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faAdd, faRemove } from '@fortawesome/free-solid-svg-icons';
import { map, Observable, startWith } from 'rxjs';

import { ScrollingModule } from '@angular/cdk/scrolling';
import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

import { Store } from '@ngrx/store';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { CompanySizeEnum } from '@app/common/models/companies.model';
import { LocalFile } from '@app/common/models/files.model';
import { JobLocationTypeEnum, JobStatusEnum, JobTypeEnum, WorkPermitEnum } from '@app/common/models/jobs.model';
import { Language, LocationCity } from '@app/common/models/location.model';
import { UserKindEnum } from '@app/common/models/users.model';
import { CompanySizePipe } from '@app/common/pipes/company-size/company-size.pipe';
import { FilterCityPipe } from '@app/common/pipes/filter-city/filter-city.pipe';
import { FilterLanguagePipe } from '@app/common/pipes/filter-language/filter-language.pipe';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';
import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobStatusPipe } from '@app/common/pipes/job-status/job-status.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { UserCompanyRoleLabelPipe } from '@app/common/pipes/user-company-role-label/user-company-role-label.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { WorkPermitPipe } from '@app/common/pipes/work-permit/work-permit.pipe';
import { selectLanguages, selectLocationCities } from '@app/common/stores/location/location.reducer';
import { fileUrl } from '@app/common/utils/file.util';
import { FileItemComponent } from '@app/ui/file-item/file-item.component';
import {
    CandidateEducationHistoryFormGroup,
    CandidateEducationHistoryFormModel,
    CandidateEmploymentHistoryFormGroup,
    CandidateEmploymentHistoryFormModel,
    CandidateSkillsFormGroup,
    CandidateSkillsFormModel,
    CandidateSpokenLanguagesFormGroup,
    CandidateSpokenLanguagesFormModel,
    UserFormCandidateFormGroup,
    UserFormCandidateFormModel,
    UserFormComponent
} from '@app/users/form/common/models/user-form.model';

const DEFAULT_PHOTO_URL = 'assets/profile-picture-default-form.png';

@Component({
    selector: 'app-user-form-candidate',
    standalone: true,
    imports: [
        CommonModule,
        ReactiveFormsModule,
        FontAwesomeModule,
        ScrollingModule,
        LetDirective,
        FormErrorMessagePipe,
        ReactiveFormsModule,
        UserCompanyRoleLabelPipe,
        UserKindLabelPipe,
        FilterCityPipe,
        JobStatusPipe,
        JobTypePipe,
        CompanySizePipe,
        JobLocationTypePipe,
        WorkPermitPipe,
        FilterLanguagePipe,
        FileItemComponent,
        IsAuthorizedPipe
    ],
    templateUrl: './user-form-candidate.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserFormCandidateComponent implements UserFormComponent, OnInit {
    @Input() singleColumn = false;

    form!: FormGroup<UserFormCandidateFormGroup>;
    spokenLanguagesForm!: FormGroup<CandidateSpokenLanguagesFormGroup>;
    skillsForm!: FormGroup<CandidateSkillsFormGroup>;
    educationHistoryForm!: FormGroup<CandidateEducationHistoryFormGroup>;
    employmentHistoryForm!: FormGroup<CandidateEmploymentHistoryFormGroup>;
    filterLocationsForm!: FormControl<string | null>;
    filterLanguageForm!: FormControl<string | null>;
    imagePreview$!: Observable<LocalFile>;
    cities$!: Observable<LocationCity[]>;
    languages$!: Observable<Language[]>;

    get formValue(): UserFormCandidateFormModel {
        const value = this.form.getRawValue();
        return {
            ...value,
            details: {
                ...value.details,
                birthDate: value.details.birthDate && new Date(value.details.birthDate).toISOString()
            },
            technical: {
                ...value.technical,
                educationHistory:
                    value.technical.educationHistory?.map((hist) => ({
                        ...hist,
                        fromDate: hist.fromDate && new Date(hist.fromDate),
                        toDate: hist.toDate && new Date(hist.toDate)
                    })) || [],
                employmentHistory:
                    value.technical.employmentHistory?.map((hist) => ({
                        ...hist,
                        fromDate: hist.fromDate && new Date(hist.fromDate),
                        toDate: hist.toDate && new Date(hist.toDate)
                    })) || []
            }
        };
    }

    get detailsForm(): UserFormCandidateFormGroup['details'] {
        return this.form.controls.details;
    }

    get jobForm(): UserFormCandidateFormGroup['job'] {
        return this.form.controls.job;
    }

    get technicalForm(): UserFormCandidateFormGroup['technical'] {
        return this.form.controls.technical;
    }

    get socialForm(): UserFormCandidateFormGroup['social'] {
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

    protected readonly jobStatuses: JobStatusEnum[] = [
        JobStatusEnum.SEARCHING,
        JobStatusEnum.OPEN_TO,
        JobStatusEnum.NOT_SEARCHING
    ];
    protected readonly jobTypes: JobTypeEnum[] = [
        JobTypeEnum.FULL_TIME,
        JobTypeEnum.PART_TIME,
        JobTypeEnum.INTERNSHIP,
        JobTypeEnum.TEMPORARY
    ];
    protected readonly jobLocationTypes: JobLocationTypeEnum[] = [
        JobLocationTypeEnum.ON_SITE,
        JobLocationTypeEnum.HYBRID,
        JobLocationTypeEnum.REMOTE
    ];
    protected readonly workPermits: WorkPermitEnum[] = [
        WorkPermitEnum.CITIZEN,
        WorkPermitEnum.PERMANENT_RESIDENT,
        WorkPermitEnum.WORK_VISA,
        WorkPermitEnum.STUDENT_VISA,
        WorkPermitEnum.TEMPORARY_RESIDENT,
        WorkPermitEnum.NO_WORK_PERMIT,
        WorkPermitEnum.OTHER
    ];
    protected readonly companySizes: CompanySizeEnum[] = [
        CompanySizeEnum.ANY,
        CompanySizeEnum.SMALL,
        CompanySizeEnum.MEDIUM,
        CompanySizeEnum.LARGE
    ];
    protected readonly userKindEnum = UserKindEnum;
    protected readonly faAdd = faAdd;
    protected readonly faRemove = faRemove;

    constructor(
        private readonly fb: FormBuilder,
        private readonly store: Store
    ) {}

    ngOnInit(): void {
        this.initForm();
        this.initSubscriptions();
    }

    onPhotoUpload(event: Event): void {
        const input = event.target as HTMLInputElement;
        const file = input.files?.[0];
        this.detailsForm.controls.photo.setValue(file || null);
        input.value = '';
    }

    onCVUpload(event: Event): void {
        const input = event.target as HTMLInputElement;
        const file = input.files?.[0];
        this.technicalForm.controls.cv.setValue(file || null);
        input.value = '';
    }

    onCVRemove(): void {
        this.technicalForm.controls.cv.setValue(null);
    }

    onAttachmentsUpload(event: Event): void {
        const input = event.target as HTMLInputElement;
        const files = input.files;
        this.technicalForm.controls.attachments.setValue(Array.from(files || []));
        input.value = '';
    }

    onAttachmentRemove(index: number): void {
        const attachments = this.technicalForm.controls.attachments.value;
        if (!attachments) {
            return;
        }
        const firstPart = attachments.slice(0, index) as File[];
        const secondPart = attachments.slice(index + 1) as File[];
        this.technicalForm.controls.attachments.setValue([...firstPart, ...secondPart]);
    }

    onVideoUpload(event: Event): void {
        const input = event.target as HTMLInputElement;
        const file = input.files?.[0];
        this.technicalForm.controls.video.setValue(file || null);
        input.value = '';
    }

    onVideoRemove(): void {
        this.technicalForm.controls.video.setValue(null);
    }

    onAddSpokenLanguage(): void {
        const control = this.spokenLanguagesForm;
        if (control.invalid) {
            return;
        }
        this.technicalForm.controls.spokenLanguages.patchValue([...this.spokenLanguages, control.getRawValue()]);
        control.markAsPristine();
        control.reset();

        this.filterLanguageForm.markAsPristine();
        this.filterLanguageForm.reset();
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

    onAddPreferredLocation(city: LocationCity): void {
        const control = this.jobForm.controls.seekLocations;
        if (control.value?.includes(city)) {
            return;
        }
        control.setValue((control.value || []).concat(city));
        control.markAsTouched();
        this.filterLocationsForm.reset();
    }

    onRemovePreferredLocation(index: number): void {
        const control = this.jobForm.controls.seekLocations;
        const value = control.value || [];
        control.setValue(value.slice(0, index).concat(value.slice(index + 1)));
        control.markAsTouched();
    }

    onSelectSpokenLanguage(language: Language): void {
        this.filterLanguageForm.setValue(language.name);
        this.spokenLanguagesForm.controls.language.setValue(language);
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
                birthDate: this.fb.control<string | null>(null, [Validators.required]),
                photo: this.fb.control<LocalFile | File | null>(null)
            }),
            job: this.fb.group({
                yearsOfExperience: this.fb.control<number | null>(null, [Validators.required, Validators.min(0)]),
                jobStatus: this.fb.control<JobStatusEnum | null>(null, [Validators.required]),
                seekJobType: this.fb.control<JobTypeEnum | null>(null),
                seekCompanySize: this.fb.control<CompanySizeEnum | null>(null),
                seekLocations: this.fb.control<LocationCity[] | null>(
                    [],
                    [Validators.required, Validators.minLength(1)]
                ),
                seekLocationType: this.fb.control<JobLocationTypeEnum | null>(null, [Validators.required]),
                seekSalary: this.fb.control<number | null>(null),
                seekValues: this.fb.control<string | null>(null),
                workPermit: this.fb.control<WorkPermitEnum | null>(null, [Validators.required]),
                noticePeriod: this.fb.control<number | null>(null)
            }),
            technical: this.fb.group({
                spokenLanguages: this.fb.control<CandidateSpokenLanguagesFormModel[] | null>([]),
                skills: this.fb.control<CandidateSkillsFormModel[] | null>([]),
                cv: this.fb.control<LocalFile | File | null>(null),
                attachments: this.fb.control<LocalFile[] | File[] | null>([]),
                video: this.fb.control<LocalFile | File | null>(null),
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
            language: this.fb.control<Language | null>(null, [Validators.required]),
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
        this.filterLocationsForm = this.fb.control<string | null>(null, [Validators.minLength(3)]);
        this.filterLanguageForm = this.fb.control<string | null>(null, [Validators.minLength(3)]);
    }

    private initSubscriptions(): void {
        this.imagePreview$ = this.detailsForm.controls.photo.valueChanges.pipe(
            startWith(this.detailsForm.controls.photo.value),
            map((file: LocalFile | File | null) => ({
                name: file?.name || '',
                url: fileUrl(file, DEFAULT_PHOTO_URL) as string
            }))
        );

        this.cities$ = this.store.select(selectLocationCities);
        this.languages$ = this.store.select(selectLanguages);
    }
}
