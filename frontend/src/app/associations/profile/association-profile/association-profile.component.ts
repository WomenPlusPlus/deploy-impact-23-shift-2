import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { AssociationProfileStore } from './association-profile.store';

@Component({
    selector: 'app-association-profile',
    standalone: true,
    imports: [CommonModule, ContentErrorComponent, ContentLoadingComponent],
    providers: [provideComponentStore(AssociationProfileStore)],
    templateUrl: './association-profile.component.html'
})
export class AssociationProfileComponent {
    id?: number;
    readonly vm$ = this.associationProfileStore.vm$;

    constructor(
        private readonly associationProfileStore: AssociationProfileStore,
        private route: ActivatedRoute,
        private router: Router
    ) {}

    ngOnInit(): void {
        this.id = Number(this.route.snapshot.paramMap.get('id'));
        if (this.id) {
            this.associationProfileStore.getProfile(this.id);
        } else {
            this.router.navigate(['/associations']);
        }
    }
}
