import { FormArray, FormControl, FormGroup } from '@angular/forms';

import { CompanySizeEnum } from '@app/common/models/companies.model';
import { JobLocationTypeEnum, JobStatusEnum, JobTypeEnum, WorkPermitEnum } from '@app/common/models/jobs.model';
import { LocationCity } from '@app/common/models/location.model';

export type CreateUserFormGroup = FormModel<CreateUserFormModel>;
export type CreateUserCandidateFormGroup = FormModel<CreateUserCandidateFormModel>;

export interface CreateUserFormModel {
    details: UserDetailsFormModel;
    social: UserSocialFormModel;
}

export interface CreateUserCandidateFormModel extends CreateUserFormModel {
    job: UserJobFormModel;
    technical: UserTechnicalFormModel;
}

export interface UserDetailsFormModel {
    firstName: string | null;
    lastName: string | null;
    preferredName: string | null;
    email: string | null;
    phoneNumber: string | null;
    birthDate: Date | null;
    photo: File | null;
}

export interface UserJobFormModel {
    yearsOfExperience: number | null;
    jobStatus: JobStatusEnum | null;
    seekJobType: JobTypeEnum | null;
    seekCompanySize: CompanySizeEnum | null;
    seekLocations: LocationCity[] | null;
    seekLocationType: JobLocationTypeEnum | null;
    seekSalary: number | null;
    seekValues: string | null;
    workPermit: WorkPermitEnum | null;
    noticePeriod: string | null;
}

export interface UserTechnicalFormModel {
    spokenLanguages: CandidateSpokenLanguagesFormModel[] | null;
    skills: CandidateSkillsFormModel[] | null;
    cv: File | null;
    attachments: File[] | null;
    video: File | null;
    educationHistory: CandidateEducationHistoryFormModel[] | null;
    employmentHistory: CandidateEmploymentHistoryFormModel[] | null;
}

export interface UserSocialFormModel {
    linkedInUrl: string | null;
    githubUrl: string | null;
    portfolioUrl: string | null;
}

export type CandidateSpokenLanguagesFormGroup = FormModel<CandidateSpokenLanguagesFormModel>;

export interface CandidateSpokenLanguagesFormModel {
    language: string | null; // TODO: enum
    level: number | null;
}

export type CandidateSkillsFormGroup = FormModel<CandidateSkillsFormModel>;

export interface CandidateSkillsFormModel {
    name: string | null; // TODO: select
    years: number | null;
}

export type CandidateEducationHistoryFormGroup = FormModel<CandidateEducationHistoryFormModel>;

export interface CandidateEducationHistoryFormModel {
    title: string | null;
    description: string | null;
    entity: string | null;
    fromDate: Date | null;
    toDate: Date | null;
    onGoing: boolean | null;
}

export type CandidateEmploymentHistoryFormGroup = FormModel<CandidateEmploymentHistoryFormModel>;

export interface CandidateEmploymentHistoryFormModel {
    title: string | null;
    description: string | null;
    company: string | null;
    fromDate: Date | null;
    toDate: Date | null;
    onGoing: boolean | null;
}

type FormModel<T> = {
    [key in keyof T]: T[key] extends (infer U | null)[]
        ? FormArray<U extends object ? FormGroup<FormModel<U>> : FormControl<U | null>>
        : T[key] extends object
        ? FormGroup<FormModel<T[key]>>
        : FormControl<T[key]>;
};
