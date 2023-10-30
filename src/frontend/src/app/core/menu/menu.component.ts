import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faSquareCaretLeft, faSquareCaretRight } from '@fortawesome/free-solid-svg-icons';
import { filter, map, Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import {
    ChangeDetectionStrategy,
    Component,
    ElementRef,
    EventEmitter,
    Input,
    NgZone,
    OnChanges,
    OnInit,
    Output
} from '@angular/core';
import { RouterModule } from '@angular/router';

import { Store } from '@ngrx/store';

import { selectProfile } from '@app/common/stores/auth/auth.reducer';
import { MenuItem } from '@app/core/menu/common/models/menu.model';
import { MENU_ITEMS_BY_KIND } from '@app/core/menu/common/utils/menu-items.util';

@Component({
    selector: 'app-menu',
    standalone: true,
    imports: [CommonModule, RouterModule, FontAwesomeModule],
    styleUrls: ['./menu.component.scss'],
    templateUrl: './menu.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class MenuComponent implements OnInit, OnChanges {
    @Input() expanded = true;
    @Input() showExpanded = true;
    @Output() expandedChange = new EventEmitter<boolean>();

    items$!: Observable<MenuItem[]>;

    protected readonly faSquareCaretLeft = faSquareCaretLeft;
    protected readonly faSquareCaretRight = faSquareCaretRight;

    constructor(
        private readonly el: ElementRef<HTMLElement>,
        private readonly zone: NgZone,
        private readonly store: Store
    ) {}

    ngOnInit(): void {
        this.initSubscriptions();
    }

    ngOnChanges(): void {
        this.zone.runOutsideAngular(() => setTimeout(() => this.checkOpenLink(), 500));
    }

    private checkOpenLink(): void {
        if (!this.expanded) {
            return;
        }

        const activeLink = this.el.nativeElement.querySelector('a.link-active');
        if (!activeLink) {
            return;
        }

        let el: Element | null = activeLink;
        while (el && el.tagName !== 'DETAILS') {
            el = el.parentElement;
        }

        if (el) {
            (el.querySelector('details summary') as HTMLDetailsElement)?.click();
        }
    }

    private initSubscriptions(): void {
        this.items$ = this.store.select(selectProfile).pipe(
            filter(Boolean),
            map(({ kind }) => MENU_ITEMS_BY_KIND[kind])
        );
    }
}
