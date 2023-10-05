import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from '@angular/core';

import { UserListItemModel } from '../../common/models/user-card.model';

@Component({
    selector: 'app-user-card',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './user-card.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserCardComponent {
    @Input() user!: UserListItemModel;
    @Input() loggedIn!: boolean;

    @Output() logIn = new EventEmitter<void>();
}
