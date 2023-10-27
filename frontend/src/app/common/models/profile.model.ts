import { AssociationProfileModel } from '@app/associations/common/models/association-profile.model';
import { UserKindEnum, UserRoleEnum, UserStateEnum } from '@app/common/models/users.model';
import { CompanyProfileModel } from '@app/companies/profile/common/models/company-profile.model';

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
    company?: CompanyProfileModel;
    association?: AssociationProfileModel;
}
