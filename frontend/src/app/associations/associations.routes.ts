import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        loadComponent: () =>
            import('./create-association/create-association.component').then((m) => m.CreateAssociationComponent)
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
            import('./edit-association/edit-association.component').then((m) => m.EditAssociationComponent)
    },
    {
        path: '',
        loadComponent: () =>
            import('./list/associations-list/associations-list.component').then((m) => m.AssociationsListComponent)
    }
];

export default routes;
