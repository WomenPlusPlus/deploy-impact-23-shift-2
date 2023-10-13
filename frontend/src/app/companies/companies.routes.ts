import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: ':id',
        loadComponent: () =>
            import('@app/companies/profile/company-profile/company-profile.component').then(
                (m) => m.CompanyProfileComponent
            )
    }
];

export default routes;
