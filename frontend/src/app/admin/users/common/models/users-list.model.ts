import { JobStatusEnum } from '@app/common/models/jobs.model';
import {
    AssociationUserRoleEnum,
    CompanyUserRoleEnum,
    UserKindEnum,
    UserStateEnum
} from '@app/common/models/users.model';

export interface UsersListModel {
    items: UsersListItemModel[];
}

export interface UsersListItemModel {
    id: number;
    kind: UserKindEnum;
    firstName: string;
    lastName: string;
    preferredName?: string;
    imageUrl: string;
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
    role: AssociationUserRoleEnum;
}

export interface UsersListCompanyModel extends UsersListItemModel {
    role: CompanyUserRoleEnum;
}

export type UsersListMode = 'short' | 'detailed';
