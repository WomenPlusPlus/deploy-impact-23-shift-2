import { UserKindEnum, UserStateEnum } from '@app/common/models/users.model';

export interface UsersList {
    items: UsersItem[];
}

export interface UsersItem {
    id: number;
    firstName: string;
    lastName: string;
    preferredName?: string;
    email: string;
    kind: UserKindEnum;
    state: UserStateEnum;
}
