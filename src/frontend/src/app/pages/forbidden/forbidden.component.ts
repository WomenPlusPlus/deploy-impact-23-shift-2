import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';

@Component({
    selector: 'app-page-not-found',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './forbidden.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ForbiddenComponent {}
