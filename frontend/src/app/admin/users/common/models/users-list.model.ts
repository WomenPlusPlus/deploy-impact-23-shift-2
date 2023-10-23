import { JobStatusEnum } from '@app/common/models/jobs.model';
import { UserRoleEnum, UserKindEnum, UserStateEnum } from '@app/common/models/users.model';

export interface UsersListModel {
    items: UsersListItemModel[];
}

export interface UsersListItemModel {
    id: number;
    kind: UserKindEnum;
    firstName: string;
    lastName: string;
    preferredName?: string;
    imageUrl?: string;
    email: string;
    state: UserStateEnum;
}

export interface UsersListCandidateModel extends UsersListItemModel {
    phoneNumber: string;
    ratingSkill: number;
    jobStatus: JobStatusEnum;
    hasCV: boolean;
    hasVideo: boolean;
}

export interface UsersListAssociationModel extends UsersListItemModel {
    role: UserRoleEnum;
}

export interface UsersListCompanyModel extends UsersListItemModel {
    role: UserRoleEnum;
}

export type UsersListMode = 'short' | 'detailed';
