import { FormControl } from '@angular/forms';

import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';

export interface CreateInviteFormGroup {
    kind: FormControl<UserKindEnum | null>;
    role: FormControl<UserRoleEnum | null>;
    email: FormControl<string | null>;
    subject: FormControl<string | null>;
    message: FormControl<string | null>;
}

export interface CreateInviteFormModel {
    kind: UserKindEnum;
    role: UserRoleEnum;
    email: string;
    subject: string;
    message: string;
}
