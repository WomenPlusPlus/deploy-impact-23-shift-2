import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';

import { SetupCandidateFormComponent } from '@app/setup/candidate/setup-candidate-form.component';

@Component({
    selector: 'app-setup-screen',
    standalone: true,
    imports: [CommonModule, SetupCandidateFormComponent],
    templateUrl: './setup-screen.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SetupScreenComponent {}
