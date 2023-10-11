import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        loadComponent: () => import('./create-company/create-company.component').then((m) => m.CreateCompanyComponent)
    }
];

export default routes;
