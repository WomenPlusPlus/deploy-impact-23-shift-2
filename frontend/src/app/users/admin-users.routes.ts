import { Routes } from '@angular/router';

import { hasUserIDGuard, isAdminGuard } from '@app/common/guards/admin.guard';

const routes: Routes = [
    {
        path: '',
        loadComponent: () => import('./list/users-list.component').then((m) => m.UsersListComponent)
    },
    {
        path: 'create',
        loadComponent: () => import('./form/create-user.component').then((m) => m.CreateUserComponent),
        canActivate: [isAdminGuard]
    },
    {
        path: ':id',
        loadComponent: () => import('./view/view-user.component').then((m) => m.ViewUserComponent),
        pathMatch: 'full'
    },
    {
        path: ':id/edit',
        loadComponent: () => import('./form/edit-user.component').then((m) => m.EditUserComponent),
        canActivate: [hasUserIDGuard]
    }
];

export default routes;
