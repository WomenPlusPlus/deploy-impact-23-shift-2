export interface CompanyProfileModel {
    id: number;
    name: string;
    address: string;
    logo: string;
    linkedin: string;
    kununu: string;
    phone: string;
    email: string;
    values: string;
    jobtypes: string;
    expectation: string;
}

export interface CompaniesListModel {
    items: CompanyProfileModel[];
}
