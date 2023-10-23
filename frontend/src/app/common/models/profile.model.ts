import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';

export interface Profile {
    id: number;
    name: string;
    email: string;
    avatar: string;
    kind: UserKindEnum;
    role?: UserRoleEnum;
}
