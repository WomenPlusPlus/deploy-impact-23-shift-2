import { FormControl } from '@angular/forms';

export interface CreateCompanyFormGroup {
    name: FormControl<string | null>;
    address: FormControl<string | null>;
    logo: FormControl<File | null>;
    linkedin: FormControl<string | null>;
    kununu: FormControl<string | null>;
    phone: FormControl<string | null>;
    email: FormControl<string | null>;
    values: FormControl<string | null>;
    jobtypes: FormControl<string | null>;
    expectation: FormControl<string | null>;
}
