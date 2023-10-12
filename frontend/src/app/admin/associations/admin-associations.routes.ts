import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        loadComponent: () =>
            import('./create-association/create-association.component').then((m) => m.CreateAssociationComponent)
    }
];

export default routes;
