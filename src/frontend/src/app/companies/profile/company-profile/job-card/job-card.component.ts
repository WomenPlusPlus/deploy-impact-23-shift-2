import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { Job } from '@app/jobs/common/models/job.model';

@Component({
    selector: 'app-job-card',
    standalone: true,
    imports: [CommonModule, RouterModule, JobLocationTypePipe, JobTypePipe],
    templateUrl: './job-card.component.html'
})
export class JobCardComponent {
    @Input() job!: Job;
}
