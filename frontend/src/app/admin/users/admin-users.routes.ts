import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        loadComponent: () => import('./list/users-list.component').then((m) => m.UsersListComponent)
    },
    {
        path: 'create',
        loadComponent: () => import('./form/create-user.component').then((m) => m.CreateUserComponent)
    }
];

export default routes;
