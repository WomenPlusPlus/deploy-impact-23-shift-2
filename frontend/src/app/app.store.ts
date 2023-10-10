import { map, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore } from '@ngrx/component-store';

interface AppState {
    menuExpanded: boolean;
}

const MENU_EXPANDED_KEY = 'shift2:menu-expanded';

@Injectable()
export class AppStore extends ComponentStore<AppState> {
    vm$ = this.select({
        menuExpanded: this.select((state) => state.menuExpanded)
    });

    constructor() {
        super({
            menuExpanded: localStorage.getItem(MENU_EXPANDED_KEY) === 'true'
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
