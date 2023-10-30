import { Routes } from '@angular/router';

import { isAdminGuard } from '@app/common/guards/admin.guard';
import { isRelatedToCompany } from '@app/common/guards/related-user.guard';

const routes: Routes = [
    {
        path: 'create',
        loadComponent: () => import('./create-company/create-company.component').then((m) => m.CreateCompanyComponent),
        canActivate: [isAdminGuard]
    },
    {
        path: ':id',
        loadComponent: () =>
            import('@app/companies/profile/company-profile/company-profile.component').then(
                (m) => m.CompanyProfileComponent
            )
    },
    {
        path: ':id/edit',
        loadComponent: () => import('./edit-company/edit-company.component').then((m) => m.EditCompanyComponent),
        canActivate: [isRelatedToCompany]
    },
    {
        path: '',
        loadComponent: () =>
            import('@app/companies/list/companies-list/companies-list.component').then((m) => m.CompaniesListComponent)
    }
];

export default routes;
