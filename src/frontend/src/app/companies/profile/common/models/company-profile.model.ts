import { LocalFile } from '@app/common/models/files.model';

export interface CompanyItem {
    id: number;
    name: string;
    logo?: LocalFile;
    linkedinUrl?: string;
    kununuUrl?: string;
    contactEmail: string;
    contactPhone: string;
    address: string;
    mission: string;
    values: string;
    jobTypes: string;
    expectation?: string;
    createdAt: string;
}

export interface CompanyList {
    items: CompanyItem[];
}
