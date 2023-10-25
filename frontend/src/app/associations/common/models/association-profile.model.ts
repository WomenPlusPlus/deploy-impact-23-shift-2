export interface AssociationProfileModel {
    id: number;
    name: string;
    logo: string;
    websiteUrl: string;
    focus: string;
}

export interface AssociationsListModel {
    items: AssociationProfileModel[];
}
