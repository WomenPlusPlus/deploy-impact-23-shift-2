import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: 'users',
        loadChildren: () => import('./users/admin-users.routes')
    },
    {
        path: 'invitations',
        loadChildren: () => import('./invitations/admin-invitations.routes')
    },
    {
        path: 'companies',
        loadChildren: () => import('./company/admin-company.routes')
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'users'
    }
];

export default routes;
