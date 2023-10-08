import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, HostBinding } from '@angular/core';

@Component({
    selector: 'app-content-loading',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './content-loading.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ContentLoadingComponent {
    @HostBinding('class') classes = 'flex items-center justify-center';
}
