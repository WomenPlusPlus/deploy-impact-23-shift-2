import { JobLocationTypeEnum, JobTypeEnum } from '@app/common/models/jobs.model';
import { LocationCity } from '@app/common/models/location.model';

export interface Job {
    benefits?: string;
    candidateOverview: string;
    company: JobCompany;
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
}

export interface JobLocation {
    city: LocationCity;
    type: JobLocationTypeEnum;
}
