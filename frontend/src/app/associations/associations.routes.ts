import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: ':id',
        loadComponent: () =>
            import('@app/associations/profile/association-profile/association-profile.component').then(
                (m) => m.AssociationProfileComponent
            )
    },
    {
        path: '',
        loadComponent: () =>
            import('@app/associations/list/associations-list/associations-list.component').then(
                (m) => m.AssociationsListComponent
            )
    }
];

export default routes;
