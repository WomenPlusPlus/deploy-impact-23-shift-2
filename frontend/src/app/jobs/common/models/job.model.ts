import { JobLocationTypeEnum, JobTypeEnum } from '@app/common/models/jobs.model';
import { LocationCity } from '@app/common/models/location.model';

export interface Job {
    id: number;
    benefits?: string;
    candidateOverview: string;
    company: JobCompany;
    creator: JobCreator;
    creationDate: string;
    experienceYearFrom?: number;
    experienceYearTo?: number;
    employmentLevelFrom?: number;
    employmentLevelTo?: number;
    jobType: JobTypeEnum;
    location: JobLocation;
    overview: string;
    rolesAndResponsibility: string;
    offerSalary?: number;
    skills: string[];
    title: string;
}

export interface JobCompany {
    id: number;
    name: string;
    mission: string;
    values: string;
    imageUrl: string;
}

export interface JobCreator {
    id: number;
    name: string;
    email: string;
    imageUrl: string;
}

export interface JobLocation {
    city: LocationCity;
    type: JobLocationTypeEnum;
}

export interface JobList {
    items: JobItem[];
}

export interface JobItem {
    id: number;
    title: string;
    jobType: JobTypeEnum;
    offerSalary: number;
    company: JobCompanyItem;
    creator: JobCreator;
    creationDate: string;
    location: JobLocation;
}

export interface JobCompanyItem {
    id: number;
    name: string;
    imageUrl: string;
}
