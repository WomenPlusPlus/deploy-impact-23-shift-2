export interface AssociationProfileModel {
    id: number;
    name: string;
    logo: string;
    url: string;
    focus: string;
}

export interface AssociationsListModel {
    items: AssociationProfileModel[];
}
