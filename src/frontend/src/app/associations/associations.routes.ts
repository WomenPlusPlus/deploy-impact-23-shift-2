import { Routes } from '@angular/router';

import { isAdminGuard } from '@app/common/guards/admin.guard';
import { isRelatedToAssociation } from '@app/common/guards/related-user.guard';

const routes: Routes = [
    {
        path: 'create',
        loadComponent: () =>
            import('./create-association/create-association.component').then((m) => m.CreateAssociationComponent),
        canActivate: [isAdminGuard]
    },
    {
        path: ':id',
        loadComponent: () =>
            import('./profile/association-profile/association-profile.component').then(
                (m) => m.AssociationProfileComponent
            )
    },
    {
        path: ':id/edit',
        loadComponent: () =>
            import('./edit-association/edit-association.component').then((m) => m.EditAssociationComponent),
        canActivate: [isRelatedToAssociation]
    },
    {
        path: '',
        loadComponent: () =>
            import('./list/associations-list/associations-list.component').then((m) => m.AssociationsListComponent)
    }
];

export default routes;
