import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { provideComponentStore } from '@ngrx/component-store';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { AssociationCardComponent } from './association-card/association-card.component';
import { AssociationsListStore } from './associations-list.store';
import { faAdd } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { RouterModule } from '@angular/router';

@Component({
    selector: 'app-associations-list',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        FontAwesomeModule,
        AssociationCardComponent,
        ContentErrorComponent,
        ContentLoadingComponent,
        FormsModule,
        LetDirective
    ],
    providers: [provideComponentStore(AssociationsListStore)],
    templateUrl: './associations-list.component.html'
})
export class AssociationsListComponent implements OnInit {
    readonly vm$ = this.associationsListStore.vm$;
    searchTerm$ = this.associationsListStore.searchString$;

    constructor(private readonly associationsListStore: AssociationsListStore) {}

    ngOnInit(): void {
        this.associationsListStore.initFilters();
        this.associationsListStore.getList();
        this.associationsListStore.initFilters();
    }

    onSearchTermChange(term: string): void {
        this.associationsListStore.updateFilterSearchTerm(term);
    }

    protected readonly faAdd = faAdd;
}
