import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterModule } from '@angular/router';

import { AppStore } from '@app/app.store';
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
export class AppComponent {
    vm$ = this.appStore.vm$;

    constructor(private readonly appStore: AppStore) {}

    onToggleMenuExpanded(): void {
        this.appStore.toggleMenuExpand();
    }
}
