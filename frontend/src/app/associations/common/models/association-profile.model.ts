export interface AssociationProfileModel {
    id: number;
    name: string;
    imageUrl: string;
    websiteUrl: string;
    focus: string;
}

export interface AssociationsListModel {
    items: AssociationProfileModel[];
}
