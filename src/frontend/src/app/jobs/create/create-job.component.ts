import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faAdd, faRemove } from '@fortawesome/free-solid-svg-icons';
import { Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { Store } from '@ngrx/store';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { JobLocationTypeEnum, JobTypeEnum } from '@app/common/models/jobs.model';
import { Language, LocationCity } from '@app/common/models/location.model';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { selectLanguages, selectLocationCities } from '@app/common/stores/location/location.reducer';
import { CreateJobState, CreateJobStore } from '@app/jobs/create/common/create-job.store';
import { CreateJobFormGroup } from '@app/jobs/create/common/models/create-job.model';
import { SelectSingleComponent } from '@app/ui/select-single/select-single.component';

@Component({
    selector: 'app-create-job',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        ReactiveFormsModule,
        FormsModule,
        FormErrorMessagePipe,
        LetDirective,
        FontAwesomeModule,
        JobLocationTypePipe,
        SelectSingleComponent,
        JobTypePipe
    ],
    providers: [CreateJobStore],
    templateUrl: './create-job.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateJobComponent implements OnInit {
    form!: FormGroup<CreateJobFormGroup>;
    skillsForm!: FormControl<string | null>;
    languagesForm!: FormControl<Language | null>;

    vm$: Observable<CreateJobState> = this.createJobStore.vm$;
    cities$!: Observable<LocationCity[]>;
    languages$!: Observable<Language[]>;

    get skills(): string[] {
        return this.form.controls.skills.value || [];
    }

    get languages(): Language[] {
        return this.form.controls.languages.value || [];
    }

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

    constructor(
        private readonly fb: FormBuilder,
        private readonly store: Store,
        private readonly createJobStore: CreateJobStore
    ) {}

    ngOnInit(): void {
        this.initForm();
        this.initSubscriptions();
    }

    onAddSkill(): void {
        const control = this.skillsForm;
        if (control.invalid) {
            return;
        }
        this.form.controls.skills.patchValue([...this.skills, control.getRawValue() || '']);
        control.markAsPristine();
        control.reset();
    }

    onRemoveSkill(index: number): void {
        const skills = this.skills;
        this.form.controls.skills.patchValue(skills.slice(0, index).concat(skills.slice(index + 1)));
    }

    onAddLanguage(): void {
        const control = this.languagesForm;
        if (control.invalid) {
            return;
        }
        this.form.controls.languages.patchValue([...this.languages, control.getRawValue() as Language]);
        control.markAsPristine();
        control.reset();
    }

    onRemoveLanguage(index: number): void {
        const languages = this.languages;
        this.form.controls.languages.patchValue(languages.slice(0, index).concat(languages.slice(index + 1)));
    }

    onSubmit(): void {
        if (!this.form.valid) {
            return;
        }
        const value = this.form.getRawValue();
        this.createJobStore.submitForm({
            title: value.title || '',
            skills: value.skills || [],
            jobType: value.jobType || '',
            salaryRangeFrom: this.rangeLeft(value.salaryRange),
            salaryRangeTo: this.rangeRight(value.salaryRange),
            experienceFrom: this.rangeLeft(value.experience),
            experienceTo: this.rangeRight(value.experience),
            benefits: value.benefits || '',
            city: value.city as LocationCity,
            languages: value.languages || [],
            locationType: value.locationType || '',
            employmentLevelFrom: this.rangeLeft(value.employmentLevel),
            employmentLevelTo: this.rangeRight(value.employmentLevel),
            candidateOverview: value.candidateOverview || '',
            overview: value.overview || '',
            rolesAndResponsibility: value.rolesAndResponsibility || '',
            startDate: value.startDate ? new Date(value.startDate).toISOString() : null
        });
    }

    private initForm(): void {
        this.form = this.fb.group({
            title: this.fb.control<string | null>(null, [Validators.required]),
            skills: this.fb.control<string[] | null>([], [Validators.required]),
            jobType: this.fb.control<string | null>(null, [Validators.required]),
            salaryRange: this.fb.control<string | null>(null, [Validators.pattern(/^\d+(-\d+)?$/)]),
            experience: this.fb.control<string | null>(null, [Validators.pattern(/^\d+(-\d+)?$/)]),
            benefits: this.fb.control<string | null>(null, [Validators.required]),
            city: this.fb.control<LocationCity | null>(null, [Validators.required]),
            languages: this.fb.control<Language[] | null>([], [Validators.required]),
            locationType: this.fb.control<string | null>(null, [Validators.required]),
            employmentLevel: this.fb.control<string | null>(null, [Validators.pattern(/^\d+(-\d+)?$/)]),
            candidateOverview: this.fb.control<string | null>(null, [Validators.required]),
            overview: this.fb.control<string | null>(null, [Validators.required]),
            rolesAndResponsibility: this.fb.control<string | null>(null, [Validators.required]),
            startDate: this.fb.control<Date | string | null>(null)
        });
        this.skillsForm = this.fb.control<string | null>(null, [Validators.required]);
        this.languagesForm = this.fb.control<Language | null>(null, [Validators.required]);
    }

    private initSubscriptions(): void {
        this.cities$ = this.store.select(selectLocationCities);
        this.languages$ = this.store.select(selectLanguages);
    }

    private rangeLeft(value: string | null): number | null {
        if (!value) {
            return null;
        }
        const ranges = value.split('-');
        const res = +ranges[0];
        if (isNaN(res)) {
            return null;
        }
        return res;
    }

    private rangeRight(value: string | null): number | null {
        if (!value) {
            return null;
        }
        const ranges = value.split('-');
        if (ranges.length < 2) {
            return null;
        }
        const res = +ranges[1];
        if (isNaN(res)) {
            return null;
        }
        return res;
    }

    protected readonly faAdd = faAdd;
    protected readonly faRemove = faRemove;
}
