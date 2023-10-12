import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';

import { CreateUserCandidateComponent } from '@app/admin/users/creation/candidate/create-user-candidate.component';

@Component({
    selector: 'app-create-user',
    standalone: true,
    imports: [CommonModule, CreateUserCandidateComponent],
    templateUrl: './create-user.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateUserComponent {}
