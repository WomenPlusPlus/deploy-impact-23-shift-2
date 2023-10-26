import { Company } from '@app/common/models/companies.model';
import { UserKindEnum, UserRoleEnum, UserStateEnum } from '@app/common/models/users.model';

export interface Profile {
    id: number;
    name: string;
    email: string;
    avatar: string;
    kind: UserKindEnum;
    role?: UserRoleEnum;
    state: UserStateEnum;
}

export interface ProfileSetup {
    inviteId: number;
    email: string;
    kind: UserKindEnum;
    role?: UserRoleEnum;
    company?: Company;
    association?: Company;
}
