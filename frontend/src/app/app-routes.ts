import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: 'admin',
        loadChildren: () => import('./admin/admin.routes')
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
