import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        loadComponent: () => import('./list/invitation-list.component').then((m) => m.InvitationListComponent)
    },
    {
        path: 'create',
        loadComponent: () => import('./create-invite/create-invite.component').then((m) => m.CreateInviteComponent)
    }
];

export default routes;
