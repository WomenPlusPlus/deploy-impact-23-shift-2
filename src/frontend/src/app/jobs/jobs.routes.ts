import { Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        loadComponent: () => import('./list/jobs-list.component').then((m) => m.JobsListComponent)
    },
    {
        path: 'create',
        loadComponent: () => import('./create/create-job.component').then((m) => m.CreateJobComponent)
    },
    {
        path: 'search',
        loadComponent: () => import('./search/search-jobs.component').then((m) => m.SearchJobsComponent)
    },
    {
        path: ':id',
        loadComponent: () => import('./details/job-details.component').then((m) => m.JobDetailsComponent)
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'users'
    }
];

export default routes;
