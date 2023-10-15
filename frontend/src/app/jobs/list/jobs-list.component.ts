import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';

@Component({
    selector: 'app-jobs-list',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './jobs-list.component.html',
    styles: [],
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class JobsListComponent {}
