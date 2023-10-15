import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';

import { DateAgoPipe } from '@app/common/pipes/date-ago/date-ago.pipe';
import { JobLocationTypePipe } from '@app/common/pipes/job-location-type/job-location-type.pipe';
import { JobTypePipe } from '@app/common/pipes/job-type/job-type.pipe';
import { JobDetailsStore } from '@app/jobs/details/job-details.store';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';

@Component({
    selector: 'app-job-details',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        JobTypePipe,
        JobLocationTypePipe,
        DateAgoPipe,
        ContentErrorComponent,
        ContentLoadingComponent
    ],
    providers: [JobDetailsStore],
    styleUrls: ['./job-details.component.scss'],
    templateUrl: './job-details.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush,
    encapsulation: ViewEncapsulation.None
})
export class JobDetailsComponent implements OnInit {
    vm$ = this.jobDetailsStore.vm$;

    constructor(
        private readonly jobDetailsStore: JobDetailsStore,
        private readonly router: Router,
        private readonly route: ActivatedRoute
    ) {}

    ngOnInit(): void {
        const id = Number(this.route.snapshot.paramMap.get('id'));
        if (id && !isNaN(id)) {
            this.jobDetailsStore.getDetails(id);
        } else {
            this.router.navigate(['/jobs']);
        }
    }
}
