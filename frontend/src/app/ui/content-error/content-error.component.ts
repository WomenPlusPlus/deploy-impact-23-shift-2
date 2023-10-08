import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, HostBinding } from '@angular/core';

@Component({
    selector: 'app-content-error',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './content-error.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ContentErrorComponent {
    @HostBinding('class') classes = 'flex items-center justify-center';
}
