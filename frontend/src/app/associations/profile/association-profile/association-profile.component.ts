import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faGlobe } from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { provideComponentStore } from '@ngrx/component-store';

import { UserKindEnum } from '@app/common/models/users.model';
import { IsAuthorizedPipe } from '@app/common/pipes/is-authorized/is-authorized.pipe';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

import { AssociationProfileStore } from './association-profile.store';

@Component({
    selector: 'app-association-profile',
    standalone: true,
    imports: [CommonModule, ContentErrorComponent, ContentLoadingComponent, FontAwesomeModule, IsAuthorizedPipe],
    providers: [provideComponentStore(AssociationProfileStore)],
    templateUrl: './association-profile.component.html'
})
export class AssociationProfileComponent {
    id?: number;
    readonly vm$ = this.associationProfileStore.vm$;

    protected readonly faGlobe = faGlobe;
    protected readonly userKindEnum = UserKindEnum;

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
