import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { Profile } from '@app/common/models/profile.model';
import { selectProfile } from '@app/common/stores/auth/auth.reducer';

@Component({
    selector: 'app-navbar',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './navbar.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class NavbarComponent implements OnInit {

    profile$!: Observable<Profile | null>;

    constructor(
        private readonly store: Store,
        protected readonly auth: AuthService
    ) {
    }

    ngOnInit(): void {
        this.initSubscriptions();
    }

    private initSubscriptions(): void {
        this.profile$ = this.store.select(selectProfile);
    }
}
