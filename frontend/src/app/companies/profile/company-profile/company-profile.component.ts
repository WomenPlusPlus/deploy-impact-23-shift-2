import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faEnvelope, faPhone, faExternalLink } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { UserKindEnum } from '@app/common/models/users.model';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { CompanyProfileStore } from './company-profile.store';
import { JobCardComponent } from './job-card/job-card.component';

@Component({
    selector: 'app-company-profile',
    standalone: true,
    imports: [
        CommonModule,
        ContentErrorComponent,
        ContentLoadingComponent,
        JobCardComponent,
        FontAwesomeModule,
        IsAuthorizedPipe,
        RouterLink
    ],
    providers: [provideComponentStore(CompanyProfileStore)],
    templateUrl: './company-profile.component.html'
})
export class CompanyProfileComponent implements OnInit {
    id?: number;
    readonly vm$ = this.companyProfileStore.vm$;

    protected readonly faEnvelope = faEnvelope;
    protected readonly faPhone = faPhone;
    protected readonly faExternalLink = faExternalLink;
    protected readonly userKindEnum = UserKindEnum;

    constructor(
        private readonly companyProfileStore: CompanyProfileStore,
        private route: ActivatedRoute,
        private router: Router
    ) {}

    ngOnInit(): void {
        this.id = Number(this.route.snapshot.paramMap.get('id'));
        if (this.id) {
            this.companyProfileStore.getProfile(this.id);
            this.companyProfileStore.getJobs(this.id);
        } else {
            this.router.navigate(['/companies']);
        }
    }
}
