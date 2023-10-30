import { LocalFile } from '@app/common/models/files.model';

export interface AssociationProfileModel {
    id: number;
    name: string;
    imageUrl?: LocalFile;
    websiteUrl: string;
    focus: string;
}

export interface AssociationsListModel {
    items: AssociationProfileModel[];
}
