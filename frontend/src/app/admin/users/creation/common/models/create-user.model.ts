import { FormArray, FormControl, FormGroup } from '@angular/forms';

import { CompanySizeEnum } from '@app/common/models/companies.model';
import { JobLocationTypeEnum, JobStatusEnum, JobTypeEnum, WorkPermitEnum } from '@app/common/models/jobs.model';
import { Language, LocationCity } from '@app/common/models/location.model';
import { UserKindEnum } from '@app/common/models/users.model';

export type CreateUserFormGroup = FormModel<CreateUserFormModel>;
export type CreateUserCandidateFormGroup = FormModel<CreateUserCandidateFormModel>;
export type CreateUserCompanyFormGroup = FormModel<CreateUserCompanyFormModel>;
export type CreateUserAssociationFormGroup = FormModel<CreateUserAssociationFormModel>;

export interface CreateUserFormModel {
    details: UserDetailsFormModel;
    social: UserSocialFormModel;
}

export interface CreateUserCandidateFormModel extends CreateUserFormModel {
    job: UserJobFormModel;
    technical: UserTechnicalFormModel;
}

export interface CreateUserCompanyFormModel extends CreateUserFormModel {
    companyId: number | null;
}

export interface CreateUserAssociationFormModel extends CreateUserFormModel {
    associationId: number | null;
}

export type CreateUserSubmissionModel = CreateUserFormModel & { kind: UserKindEnum };

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
    noticePeriod: number | null;
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
    language: Language | null;
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

export interface CreateUserResponse {
    id: number;
    userId: number;
}

type FormModel<T> = {
    [key in keyof T]: T[key] extends (infer U | null)[]
        ? FormArray<U extends object ? FormGroup<FormModel<U>> : FormControl<U | null>>
        : T[key] extends object
        ? FormGroup<FormModel<T[key]>>
        : FormControl<T[key]>;
};
