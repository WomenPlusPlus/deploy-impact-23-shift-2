import { FormControl } from '@angular/forms';

import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';

export interface CreateInviteFormGroup {
    kind: FormControl<UserKindEnum | null>;
    role: FormControl<UserRoleEnum | null>;
    companyId: FormControl<number | null>;
    associationId: FormControl<number | null>;
    email: FormControl<string | null>;
    subject: FormControl<string | null>;
    message: FormControl<string | null>;
}

export interface CreateInviteFormModel {
    kind: UserKindEnum | null;
    role: UserRoleEnum | null;
    email: string | null;
    subject: string | null;
    message: string | null;
}
