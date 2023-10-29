import { JobStatusEnum } from '@app/common/models/jobs.model';
import { UserKindEnum, UserRoleEnum, UserStateEnum } from '@app/common/models/users.model';

export interface UsersList {
    items: UsersItem[];
}

export interface UsersItem {
    id: number;
    kind: UserKindEnum;
    role?: UserRoleEnum;
    associationId?: number;
    companyId?: number;
    firstName: string;
    lastName: string;
    preferredName?: string;
    imageUrl?: string;
    email: string;
    state: UserStateEnum;
}

export interface UsersCandidateItem extends UsersItem {
    phoneNumber: string;
    ratingSkill: number;
    jobStatus: JobStatusEnum;
    hasCV: boolean;
    hasVideo: boolean;
}

export type UsersListMode = 'short' | 'detailed';
