import { FormControl } from '@angular/forms';

export interface CreateAssociationFormGroup {
    name: FormControl<string | null>;
    logo: FormControl<File | null>;
    websiteUrl: FormControl<string | null>;
    focus: FormControl<string | null>;
}
