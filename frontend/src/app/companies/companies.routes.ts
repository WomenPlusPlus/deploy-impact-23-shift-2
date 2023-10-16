import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: ':id',
        loadComponent: () =>
            import('@app/companies/profile/company-profile/company-profile.component').then(
                (m) => m.CompanyProfileComponent
            )
    },
    {
        path: '',
        loadComponent: () =>
            import('@app/companies/list/companies-list/companies-list.component').then((m) => m.CompaniesListComponent)
    }
];

export default routes;
