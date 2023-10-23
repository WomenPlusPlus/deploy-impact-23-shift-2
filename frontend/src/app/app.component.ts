import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { RouterModule } from '@angular/router';

import { Store } from '@ngrx/store';

import { AppStore } from '@app/app.store';
import { LocationActions } from '@app/common/stores/location/location.actions';
import { MenuComponent } from '@app/core/menu/menu.component';
import { NavbarComponent } from '@app/core/navbar/navbar.component';

@Component({
    selector: 'app-root',
    standalone: true,
    imports: [CommonModule, RouterModule, MenuComponent, NavbarComponent],
    providers: [AppStore],
    templateUrl: './app.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class AppComponent implements OnInit {
    vm$ = this.appStore.vm$;

    constructor(
        private readonly store: Store,
        private readonly appStore: AppStore
    ) {
    }

    ngOnInit(): void {
        this.store.dispatch(LocationActions.loadInitialData());
    }

    onToggleMenuExpanded(): void {
        this.appStore.toggleMenuExpand();
    }
}
