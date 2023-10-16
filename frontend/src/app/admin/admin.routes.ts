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
        path: 'companies/:id',
        loadComponent: () =>
            import('@app/companies/profile/company-profile/company-profile.component').then(
                (m) => m.CompanyProfileComponent
            )
    },
    {
        path: 'companies',
        loadChildren: () => import('./company/admin-company.routes')
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
