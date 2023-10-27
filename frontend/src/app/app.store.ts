import { map, Observable, switchMap, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore } from '@ngrx/component-store';
import { Store } from '@ngrx/store';

import { selectProfile } from '@app/common/stores/auth/auth.reducer';
import { isAuthenticated$ } from '@app/common/utils/auth.util';

interface AppState {
    menuExpanded: boolean;
}

const MENU_EXPANDED_KEY = 'shift2:menu-expanded';

@Injectable()
export class AppStore extends ComponentStore<AppState> {
    private authenticated$ = isAuthenticated$(this.store);

    vm$ = this.select({
        authenticated: this.authenticated$,
        validated: this.authenticated$.pipe(
            switchMap((authenticated) =>
                this.store.select(selectProfile).pipe(map((profile) => authenticated && profile?.state === 'ACTIVE'))
            )
        ),
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
