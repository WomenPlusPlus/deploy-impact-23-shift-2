import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: ':id',
        loadComponent: () =>
            import('@app/associations/profile/association-profile/association-profile.component').then(
                (m) => m.AssociationProfileComponent
            )
    }
];

export default routes;
