import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';

export { DescriptionListItemComponent } from '@app/ui/description-list/description-list-item/description-list-item.component';

@Component({
    selector: 'app-description-list',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './description-list.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class DescriptionListComponent {}
