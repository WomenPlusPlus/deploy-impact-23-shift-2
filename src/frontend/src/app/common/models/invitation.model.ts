import { UserKindEnum, UserRoleEnum } from '@app/common/models/users.model';

export interface InvitationList {
    items: InviteItem[];
}

export interface InviteItem {
    id: number;
    creatorId: number;
    kind: UserKindEnum;
    role?: UserRoleEnum;
    entityId?: number;
    email: string;
    state: InviteStateEnum;
    expireAt: string;
    createdAt: string;
}

export enum InviteStateEnum {
    CREATED = 'CREATED',
    PENDING = 'PENDING',
    ERROR = 'ERROR',
    ACCEPTED = 'ACCEPTED',
    CANCELLED = 'CANCELLED'
}
