import { ChangeDetectionStrategy, Component } from '@angular/core';

@Component({
    selector: 'app-temporary',
    templateUrl: './temporary.component.html',
    styleUrls: ['./temporary.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class TemporaryComponent {}
