import { map, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore } from '@ngrx/component-store';
import { Store } from '@ngrx/store';
import { isAuthenticated$ } from '@app/common/utils/auth.util';

interface AppState {
    menuExpanded: boolean;
}

const MENU_EXPANDED_KEY = 'shift2:menu-expanded';

@Injectable()
export class AppStore extends ComponentStore<AppState> {
    vm$ = this.select({
        authenticated: isAuthenticated$(this.store),
        menuExpanded: this.select((state) => state.menuExpanded)
    });

    constructor(private readonly store: Store) {
        super({
            menuExpanded: localStorage.getItem(MENU_EXPANDED_KEY) !== 'false'
        });
    }

    toggleMenuExpand = this.effect((trigger: Observable<void>) =>
        trigger.pipe(
            map(() => !this.state().menuExpanded),
            tap((menuExpanded) => this.patchState({ menuExpanded })),
            tap((menuExpanded) => localStorage.setItem(MENU_EXPANDED_KEY, `${menuExpanded}`))
        )
    );
}
