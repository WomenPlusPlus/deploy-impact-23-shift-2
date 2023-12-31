import { FormArray, FormControl, FormGroup } from '@angular/forms';

import { CompanySizeEnum } from '@app/common/models/companies.model';
import { LocalFile } from '@app/common/models/files.model';
import { JobLocationTypeEnum, JobStatusEnum, JobTypeEnum, WorkPermitEnum } from '@app/common/models/jobs.model';
import { Language, LocationCity } from '@app/common/models/location.model';
import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';

export interface UserFormComponent<T extends UserFormGroup = any, S extends UserFormModel = any> {
    form: FormGroup<T>;
    formValue: S;
}

export type UserFormGroup = FormModel<UserFormModel>;
export type UserFormCandidateFormGroup = FormModel<UserFormCandidateFormModel>;
export type UserFormCompanyFormGroup = FormModel<UserFormCompanyFormModel>;
export type UserFormAssociationFormGroup = FormModel<UserFormAssociationFormModel>;

export interface UserFormModel {
    details: UserDetailsFormModel;
    social: UserSocialFormModel;
}

export interface UserFormCandidateFormModel extends UserFormModel {
    job: UserJobFormModel;
    technical: UserTechnicalFormModel;
}

export interface UserFormCompanyFormModel extends UserFormModel {
    companyId: number | null;
}

export interface UserFormAssociationFormModel extends UserFormModel {
    associationId: number | null;
}

export type UserFormSubmissionModel = UserFormModel & { kind: UserKindEnum; role?: UserRoleEnum };

export interface UserDetailsFormModel {
    firstName: string | null;
    lastName: string | null;
    preferredName: string | null;
    email: string | null;
    phoneNumber: string | null;
    birthDate: string | null;
    photo: LocalFile | File | null;
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
    cv: LocalFile | File | null;
    attachments: LocalFile[] | File[] | null;
    video: LocalFile | File | null;
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
    name: string | null;
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
