import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input } from '@angular/core';

@Component({
    selector: 'app-description-list app-item',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './description-list-item.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class DescriptionListItemComponent {
    @Input() size: 'sm' | 'md' | 'lg' = 'sm'; // this can be extended to other sizes
}
