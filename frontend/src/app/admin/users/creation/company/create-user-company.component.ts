import { CdkFixedSizeVirtualScroll, CdkVirtualForOf, CdkVirtualScrollViewport } from '@angular/cdk/scrolling';
import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';

import {
    CreateUserCompanyFormGroup,
    CreateUserCompanyFormModel
} from '@app/admin/users/creation/common/models/create-user.model';
import { CreateUserFormComponent } from '@app/admin/users/creation/create-user.component';
import { CreateUserStore } from '@app/admin/users/creation/create-user.store';
import { CreateUserGenericComponent } from '@app/admin/users/creation/generic/create-user-generic.component';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { Company } from '@app/common/models/companies.model';
import { FilterFusePipe } from '@app/common/pipes/filter-fuse/filter-fuse.pipe';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { SelectSingleComponent } from '@app/ui/select-single/select-single.component';

@Component({
    selector: 'app-create-user-company',
    standalone: true,
    imports: [
        CommonModule,
        ReactiveFormsModule,
        FormsModule,
        CdkFixedSizeVirtualScroll,
        CdkVirtualForOf,
        CdkVirtualScrollViewport,
        LetDirective,
        CreateUserGenericComponent,
        FilterFusePipe,
        FormErrorMessagePipe,
        SelectSingleComponent
    ],
    templateUrl: './create-user-company.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateUserCompanyComponent implements CreateUserFormComponent, OnInit {
    @ViewChild(CreateUserGenericComponent, { static: true })
    childFormComponent!: CreateUserGenericComponent;

    companyIdForm!: FormControl<number | null>;
    searchCompanyForm!: FormControl<string | null>;
    companies$ = this.createUserStore.companies$;

    get form(): FormGroup<CreateUserCompanyFormGroup> {
        return this.fb.group({
            ...this.childFormComponent.form.controls,
            companyId: this.companyIdForm
        });
    }

    get formValue(): CreateUserCompanyFormModel {
        return {
            ...this.childFormComponent.formValue,
            companyId: this.companyIdForm.value
        };
    }

    constructor(
        private readonly fb: FormBuilder,
        private readonly createUserStore: CreateUserStore
    ) {}

    ngOnInit(): void {
        this.loadData();
        this.initForm();
    }

    onSelectCompany(company: Company): void {
        const control = this.companyIdForm;
        control.markAsTouched();
        control.setValue(company.id);

        this.searchCompanyForm.setValue(company.name);
    }

    onDeselectCompany(): void {
        const control = this.companyIdForm;
        if (control.value) {
            control.markAsTouched();
            control.setValue(null);
        }
    }

    private loadData(): void {
        this.createUserStore.loadCompanies();
    }

    private initForm(): void {
        this.companyIdForm = this.fb.control<number | null>(null, [Validators.required]);
        this.searchCompanyForm = this.fb.control<string | null>(null, [Validators.minLength(3)]);
    }
}
