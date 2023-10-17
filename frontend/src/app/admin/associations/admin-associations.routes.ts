import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        loadComponent: () =>
            import('./create-association/create-association.component').then((m) => m.CreateAssociationComponent)
    },
    {
        path: ':id/edit',
        loadComponent: () =>
            import('./edit-association/edit-association.component').then((m) => m.EditAssociationComponent)
    }
];

export default routes;
