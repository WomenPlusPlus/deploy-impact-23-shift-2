import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: 'create',
        loadComponent: () => import('./create-company/create-company.component').then((m) => m.CreateCompanyComponent)
    },
    {
        path: ':id/edit',
        loadComponent: () => import('./edit-company/edit-company.component').then((m) => m.EditCompanyComponent)
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'create'
    }
];

export default routes;
