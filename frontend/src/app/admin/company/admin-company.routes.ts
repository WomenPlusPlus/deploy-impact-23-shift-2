import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        loadComponent: () => import('./create-company/create-company.component').then((m) => m.CreateCompanyComponent)
    },
    {
        path: ':id/edit',
        loadComponent: () => import('./edit-company/edit-company.component').then((m) => m.EditCompanyComponent)
    }
];

export default routes;
