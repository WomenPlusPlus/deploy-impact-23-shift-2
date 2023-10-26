import { Routes } from '@angular/router';

import { authenticatedGuard, notAuthenticatedGuard } from '@app/common/guards/authenticated.guard';

const routes: Routes = [
    {
        path: 'login',
        loadComponent: () => import('./pages/login/login.component').then((m) => m.LoginComponent),
        canActivate: [notAuthenticatedGuard]
    },
    {
        path: 'admin',
        loadChildren: () => import('./admin/admin.routes'),
        canActivate: [authenticatedGuard]
    },
    {
        path: 'companies',
        loadChildren: () => import('./companies/companies.routes'),
        canActivate: [authenticatedGuard]
    },
    {
        path: 'jobs',
        loadChildren: () => import('./jobs/jobs.routes'),
        canActivate: [authenticatedGuard]
    },
    {
        path: 'associations',
        loadChildren: () => import('./associations/associations.routes'),
        canActivate: [authenticatedGuard]
    },
    {
        path: 'dashboard',
        loadChildren: () => import('./dashboard/dashboard.routes'),
        canActivate: [authenticatedGuard]
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'admin'
    },
    {
        path: '**',
        pathMatch: 'full',
        loadComponent: () =>
            import('./pages/page-not-found/page-not-found.component').then((m) => m.PageNotFoundComponent)
    }
];

export default routes;
