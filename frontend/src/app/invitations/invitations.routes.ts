import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: 'create',
        loadComponent: () => import('./create-invite/create-invite.component').then((m) => m.CreateInviteComponent)
    },
    {
        path: '',
        loadComponent: () => import('./list/invitation-list.component').then((m) => m.InvitationListComponent)
    }
];

export default routes;
