import { CompanySizeEnum } from '@app/common/models/companies.model';
import { LocalFile } from '@app/common/models/files.model';
import { JobLocationTypeEnum, JobStatusEnum, JobTypeEnum, WorkPermitEnum } from '@app/common/models/jobs.model';
import { Language } from '@app/common/models/location.model';

export enum UserKindEnum {
    ADMIN = 'ADMIN',
    ASSOCIATION = 'ASSOCIATION',
    COMPANY = 'COMPANY',
    CANDIDATE = 'CANDIDATE'
}

export enum UserRoleEnum {
    ADMIN = 'ADMIN',
    USER = 'USER'
}

export enum UserStateEnum {
    ACTIVE = 'ACTIVE',
    ANONYMOUS = 'ANONYMOUS',
    DELETED = 'DELETED'
}

export interface UserDetails {
    id: number;
    kind: UserKindEnum;
    firstName: string;
    lastName: string;
    preferredName: string;
    email: string;
    phoneNumber: string;
    birthDate: string;
    photo: LocalFile | null;
    linkedInUrl: string;
    githubUrl: string;
    portfolioUrl: string;
}

export interface AssociationUserDetails extends UserDetails {
    associationUserId: number;
    associationId: number;
    role: UserRoleEnum;
}

export interface CandidateDetails extends UserDetails {
    candidateId: number;
    yearsOfExperience: number;
    jobStatus: JobStatusEnum;
    seekJobType: JobTypeEnum;
    seekCompanySize: CompanySizeEnum;
    seekLocations: CandidateSeekLocation[];
    seekLocationType: JobLocationTypeEnum;
    seekSalary: number;
    seekValues: string;
    workPermit: WorkPermitEnum;
    noticePeriod: number;
    spokenLanguages: CandidateSpokenLanguage[];
    skills: CandidateSkill[];
    cv: LocalFile | null;
    attachments: LocalFile[];
    video: LocalFile | null;
    educationHistory: CandidateEducationHistory[];
    employmentHistory: CandidateEmploymentHistory[];
}

export interface CompanyUserDetails extends UserDetails {
    companyUserId: number;
    companyId: number;
    role: UserRoleEnum;
}

export interface CandidateSeekLocation {
    id: number;
    name: string;
}

export interface CandidateSpokenLanguage {
    language: Language;
    level: number;
}

export interface CandidateSkill {
    name: string;
    years: number;
}

export interface CandidateEducationHistory {
    title: string;
    description: string;
    entity: string;
    fromDate: string;
    toDate: string;
}

export interface CandidateEmploymentHistory {
    title: string;
    description: string;
    company: string;
    fromDate: string;
    toDate: string;
}
