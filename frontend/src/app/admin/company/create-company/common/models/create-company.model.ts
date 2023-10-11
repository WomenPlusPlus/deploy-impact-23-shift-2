import { FormControl } from '@angular/forms';

export interface CreateCompanyFormGroup {
    name: FormControl<string | null>;
    address: FormControl<string | null>;
    logo: FormControl<string | null>;
    linkedin: FormControl<string | null>;
    kununu: FormControl<string | null>;
    phone: FormControl<string | null>;
    email: FormControl<string | null>;
    values: FormControl<string | null>;
    jobtypes: FormControl<string | null>;
    expectation: FormControl<string | null>;
}

export interface CreateCompanyFormModel {
    name: string | null;
    address: string | null;
    logo: string | null;
    linkedin: string | null;
    kununu: string | null;
    phone: string | null;
    email: string | null;
    values: string | null;
    jobtypes: string | null;
    expectation: string | null;
}
