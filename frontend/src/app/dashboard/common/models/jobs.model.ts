import { JobLocationTypeEnum, JobTypeEnum } from '@app/common/models/jobs.model';
import { LocationCity } from '@app/common/models/location.model';

export interface JobListings {
    items: JobListing[];
}

export interface JobListing {
    id: number;
    title: string;
    jobType: JobTypeEnum;
    creator: JobCreator;
    creationDate: string;
    location: JobLocation;
}

export interface JobCreator {
    id: number;
    name: string;
    email: string;
}

export interface JobLocation {
    city: LocationCity;
    type: JobLocationTypeEnum;
}
