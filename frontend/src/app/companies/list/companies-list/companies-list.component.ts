import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faAdd } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { CompaniesListStore } from './companies-list.store';
import { CompanyCardComponent } from './company-card/company-card.component';

@Component({
    selector: 'app-companies-list',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        FontAwesomeModule,
        CompanyCardComponent,
        ContentErrorComponent,
        ContentLoadingComponent,
        FormsModule,
        LetDirective
    ],
    providers: [provideComponentStore(CompaniesListStore)],
    templateUrl: './companies-list.component.html'
})
export class CompaniesListComponent implements OnInit {
    readonly vm$ = this.companiesListStore.vm$;
    searchTerm$ = this.companiesListStore.searchString$;

    protected readonly faAdd = faAdd;

    constructor(private readonly companiesListStore: CompaniesListStore) {}

    ngOnInit(): void {
        this.companiesListStore.initFilters();
        this.companiesListStore.getList();
        this.companiesListStore.initFilters();
    }

    onSearchTermChange(term: string): void {
        this.companiesListStore.updateFilterSearchTerm(term);
    }
}
