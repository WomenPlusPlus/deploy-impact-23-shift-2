import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: 'login',
        loadComponent: () =>
            import('./pages/login/login.component').then((m) => m.LoginComponent)
    },
    {
        path: 'admin',
        loadChildren: () => import('./admin/admin.routes')
    },
    {
        path: 'companies',
        loadChildren: () => import('./companies/companies.routes')
    },
    {
        path: 'jobs',
        loadChildren: () => import('./jobs/jobs.routes')
    },
    {
        path: 'associations',
        loadChildren: () => import('./associations/associations.routes')
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'login'
    },
    {
        path: '**',
        pathMatch: 'full',
        loadComponent: () =>
            import('./pages/page-not-found/page-not-found.component').then((m) => m.PageNotFoundComponent)
    }
];

export default routes;
