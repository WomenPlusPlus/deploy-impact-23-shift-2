import { FormControl } from '@angular/forms';

import { Language, LocationCity } from '@app/common/models/location.model';

export interface CreateJobFormGroup {
    title: FormControl<string | null>;
    skills: FormControl<string[] | null>;
    jobType: FormControl<string | null>;
    salaryRange: FormControl<string | null>;
    experience: FormControl<string | null>;
    benefits: FormControl<string | null>;
    city: FormControl<LocationCity | null>;
    languages: FormControl<Language[] | null>;
    locationType: FormControl<string | null>;
    employmentLevel: FormControl<string | null>;
    candidateOverview: FormControl<string | null>;
    overview: FormControl<string | null>;
    rolesAndResponsibility: FormControl<string | null>;
    startDate: FormControl<Date | string | null>;
}

export interface CreateJobFormModel {
    title: string;
    skills: string[];
    jobType: string;
    salaryRangeFrom: number | null;
    salaryRangeTo: number | null;
    experienceFrom: number | null;
    experienceTo: number | null;
    benefits: string;
    city: LocationCity;
    languages: Language[];
    locationType: string;
    employmentLevelFrom: number | null;
    employmentLevelTo: number | null;
    candidateOverview: string;
    overview: string;
    rolesAndResponsibility: string;
    startDate: string | null;
}
