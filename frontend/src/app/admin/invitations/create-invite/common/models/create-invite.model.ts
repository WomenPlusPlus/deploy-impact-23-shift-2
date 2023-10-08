import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';

export interface CreateInviteFormModel {
    kind: UserKindEnum;
    role: UserRoleEnum;
    email: string;
    subject: string;
    message: string;
}
