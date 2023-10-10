import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import {
    faBrain,
    faBuilding,
    faCircleInfo,
    faEnvelope,
    faList,
    faListCheck,
    faSitemap,
    faSquareCaretLeft,
    faSquareCaretRight,
    faUser
} from '@fortawesome/free-solid-svg-icons';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from '@angular/core';
import { RouterModule } from '@angular/router';

@Component({
    selector: 'app-menu',
    standalone: true,
    imports: [CommonModule, RouterModule, FontAwesomeModule],
    templateUrl: './menu.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class MenuComponent {
    @Input() expanded = true;
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
}
