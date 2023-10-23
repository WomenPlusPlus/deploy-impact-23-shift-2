import { ChangeDetectionStrategy, Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CandidateDetails, UserDetails } from '@app/common/models/users.model';
import { faFile, faFileAlt, faPlay } from '@fortawesome/free-solid-svg-icons';
import { CompanySizePipe } from '@app/common/pipes/company-size/company-size.pipe';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobStatusPipe } from '@app/common/pipes/job-status/job-status.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { WorkPermitPipe } from '@app/common/pipes/work-permit/work-permit.pipe';

@Component({
    selector: 'app-view-candidate-sections',
    standalone: true,
    imports: [CommonModule, CompanySizePipe, FontAwesomeModule, JobLocationTypePipe, JobStatusPipe, JobTypePipe, WorkPermitPipe],
    templateUrl: './view-candidate-sections.component.html',
    styles: [
        `
            :host {
                display: contents;
            }
        `
    ],
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ViewCandidateSectionsComponent {
    @Input() user!: UserDetails;

    get candidate(): CandidateDetails {
        return this.user as CandidateDetails;
    }

    protected readonly faFileAlt = faFileAlt;
    protected readonly faFile = faFile;
    protected readonly faPlay = faPlay;
}
