import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';

import { provideComponentStore } from '@ngrx/component-store';

import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { CompaniesListStore } from './companies-list.store';
import { CompanyCardComponent } from './company-card/company-card.component';

@Component({
    selector: 'app-companies-list',
    standalone: true,
    imports: [CommonModule, CompanyCardComponent, ContentErrorComponent, ContentLoadingComponent],
    providers: [provideComponentStore(CompaniesListStore)],
    templateUrl: './companies-list.component.html'
})
export class CompaniesListComponent implements OnInit {
    readonly vm$ = this.companiesListStore.vm$;

    constructor(private readonly companiesListStore: CompaniesListStore) {}

    ngOnInit(): void {
        this.companiesListStore.getList();
    }
}
