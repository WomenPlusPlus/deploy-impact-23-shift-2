import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterModule } from '@angular/router';

import { CompanyProfileModel } from '@app/companies/common/models/company-profile.model';

@Component({
    selector: 'app-company-card',
    standalone: true,
    imports: [CommonModule, RouterModule],
    templateUrl: './company-card.component.html'
})
export class CompanyCardComponent {
    @Input() company!: CompanyProfileModel;
}
