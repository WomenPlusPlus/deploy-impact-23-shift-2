import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import {
    faBrain,
    faBuilding,
    faCircleInfo,
    faEnvelope,
    faList,
    faListCheck,
    faSheetPlastic,
    faSitemap,
    faSquareCaretLeft,
    faSquareCaretRight,
    faUser,
    faChartLine
} from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import {
    ChangeDetectionStrategy,
    Component,
    ElementRef,
    EventEmitter,
    Input,
    NgZone,
    OnChanges,
    Output
} from '@angular/core';
import { RouterModule } from '@angular/router';

@Component({
    selector: 'app-menu',
    standalone: true,
    imports: [CommonModule, RouterModule, FontAwesomeModule],
    styleUrls: ['./menu.component.scss'],
    templateUrl: './menu.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class MenuComponent implements OnChanges {
    @Input() expanded = true;
    @Input() showExpanded = true;
    @Output() expandedChange = new EventEmitter<boolean>();

    protected readonly faEnvelope = faEnvelope;
    protected readonly faBuilding = faBuilding;
    protected readonly faSitemap = faSitemap;
    protected readonly faUser = faUser;
    protected readonly faBrain = faBrain;
    protected readonly faCircleInfo = faCircleInfo;
    protected readonly faListCheck = faListCheck;
    protected readonly faList = faList;
    protected readonly faSquareCaretLeft = faSquareCaretLeft;
    protected readonly faSquareCaretRight = faSquareCaretRight;
    protected readonly faSheetPlastic = faSheetPlastic;
    protected readonly faChartLine = faChartLine;

    constructor(
        private readonly el: ElementRef<HTMLElement>,
        private readonly zone: NgZone
    ) {}

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
}
