import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: 'invitations',
        loadChildren: () => import('./invitations/admin-invitations.routes')
    },
    {
        path: 'companies',
        loadChildren: () => import('./company/admin-company.routes')
    },
    {
        path: 'associations/:id',
        loadComponent: () =>
            import('@app/associations/profile/association-profile/association-profile.component').then(
                (m) => m.AssociationProfileComponent
            )
    },
    {
        path: 'associations',
        loadChildren: () => import('./associations/admin-associations.routes')
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'users'
    }
];

export default routes;
