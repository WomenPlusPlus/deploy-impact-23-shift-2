export interface UsersList {
    items: UsersItem[];
}

export interface UsersItem {
    id: number;
    firstName: string;
    lastName: string;
    preferredName?: string;
    email: string;
}
