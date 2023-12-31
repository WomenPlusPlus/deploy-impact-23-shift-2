import { Routes } from '@angular/router';

import { authenticatedGuard, notAuthenticatedGuard } from '@app/common/guards/authenticated.guard';
import { activatedGuard, invitedGuard } from '@app/common/guards/setup.guard';

const routes: Routes = [
    {
        path: 'login',
        loadComponent: () => import('./pages/login/login.component').then((m) => m.LoginComponent),
        canActivate: [notAuthenticatedGuard]
    },
    {
        path: 'setup',
        loadComponent: () => import('./setup/setup-screen.component').then((m) => m.SetupScreenComponent),
        canActivate: [authenticatedGuard, invitedGuard]
    },
    {
        path: 'users',
        loadChildren: () => import('./users/admin-users.routes'),
        canActivate: [authenticatedGuard, activatedGuard]
    },
    {
        path: 'invitations',
        loadChildren: () => import('./invitations/invitations.routes')
    },
    {
        path: 'associations',
        loadChildren: () => import('./associations/associations.routes')
    },
    {
        path: 'companies',
        loadChildren: () => import('./companies/companies.routes'),
        canActivate: [authenticatedGuard, activatedGuard]
    },
    {
        path: 'jobs',
        loadChildren: () => import('./jobs/jobs.routes'),
        canActivate: [authenticatedGuard, activatedGuard]
    },
    {
        path: 'associations',
        loadChildren: () => import('./associations/associations.routes'),
        canActivate: [authenticatedGuard, activatedGuard]
    },
    {
        path: 'dashboard',
        loadChildren: () => import('./dashboard/dashboard.routes'),
        canActivate: [authenticatedGuard, activatedGuard]
    },
    {
        path: 'forbidden',
        loadComponent: () => import('./pages/forbidden/forbidden.component').then((m) => m.ForbiddenComponent)
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'dashboard'
    },
    {
        path: '**',
        pathMatch: 'full',
        loadComponent: () =>
            import('./pages/page-not-found/page-not-found.component').then((m) => m.PageNotFoundComponent)
    }
];

export default routes;
