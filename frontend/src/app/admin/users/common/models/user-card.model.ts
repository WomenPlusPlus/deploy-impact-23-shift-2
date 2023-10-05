export interface UserListModel {
    items: UserListItemModel[];
}

export interface UserListItemModel {
    title: string;
    content: string;
    imageUrl: string;
}
