import { Routes } from '@angular/router';

const routes: Routes = [
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
