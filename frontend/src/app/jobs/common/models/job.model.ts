import { JobLocationTypeEnum, JobTypeEnum } from '@app/common/models/jobs.model';
import { LocationCity } from '@app/common/models/location.model';

export interface Job {
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
