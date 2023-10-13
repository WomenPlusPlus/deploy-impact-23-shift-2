import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { CompanyProfileStore } from './company-profile.store';

@Component({
    selector: 'app-company-profile',
    standalone: true,
    imports: [CommonModule, ContentErrorComponent, ContentLoadingComponent],
    providers: [provideComponentStore(CompanyProfileStore)],
    templateUrl: './company-profile.component.html'
})
export class CompanyProfileComponent implements OnInit {
    id?: number;
    readonly vm$ = this.companyProfileStore.vm$;

    constructor(
        private readonly companyProfileStore: CompanyProfileStore,
        private route: ActivatedRoute
    ) {}

    ngOnInit(): void {
        this.id = Number(this.route.snapshot.paramMap.get('id'));
        if (this.id) {
            this.companyProfileStore.getProfile(this.id);
        }
    }
}
