import { FormArray, FormControl, FormGroup } from '@angular/forms';

export type CreateUserCandidateFormGroup = FormModel<CreateUserCandidateFormModel>;

export interface CreateUserCandidateFormModel {
    details: {
        firstName: string | null;
        lastName: string | null;
        preferredName: string | null;
        email: string | null;
        phoneNumber: string | null;
        birthDate: Date | null;
        photo: File | null;
    };
    job: {
        yearsOfExperience: number | null;
        jobStatus: string | null; // TODO: enum
        seekJobType: string | null; // TODO: enum
        seekCompanySize: string | null; // TODO: enum
        seekLocations: string[] | null; // TODO: location service?
        seekLocationType: string | null; // TODO: enum
        seekSalary: number | null;
        seekValues: string | null;
        workPermit: string | null; // TODO: enum
        noticePeriod: string | null;
    };
    technical: {
        spokenLanguages: CandidateSpokenLanguagesFormModel[] | null;
        skills: CandidateSkillsFormModel[] | null;
        cv: File | null;
        attachments: File[] | null;
        video: File | null;
        educationHistory: CandidateEducationHistoryFormModel[] | null;
        employmentHistory: CandidateEmploymentHistoryFormModel[] | null;
    };
    social: {
        linkedInUrl: string | null;
        githubUrl: string | null;
        portfolioUrl: string | null;
    };
}

export type CandidateSpokenLanguagesFormGroup = FormModel<CandidateSpokenLanguagesFormModel>;

export interface CandidateSpokenLanguagesFormModel {
    language: string | null; // TODO: enum
    level: number | null;
}

export type CandidateSkillsFormGroup = FormModel<CandidateSkillsFormModel>;

export interface CandidateSkillsFormModel {
    name: string | null; // TODO: enum
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
