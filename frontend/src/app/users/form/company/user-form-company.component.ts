import { CdkFixedSizeVirtualScroll, CdkVirtualForOf, CdkVirtualScrollViewport } from '@angular/cdk/scrolling';
import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { FilterFusePipe } from '@app/common/pipes/filter-fuse/filter-fuse.pipe';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { SelectSingleComponent } from '@app/ui/select-single/select-single.component';
import {
    UserFormCompanyFormGroup,
    UserFormCompanyFormModel,
    UserFormComponent
} from '@app/users/form/common/models/user-form.model';
import { UserFormGenericComponent } from '@app/users/form/generic/user-form-generic.component';
import { UserFormStore } from '@app/users/form/user-form.store';

@Component({
    selector: 'app-user-form-company',
    standalone: true,
    imports: [
        CommonModule,
        ReactiveFormsModule,
        FormsModule,
        CdkFixedSizeVirtualScroll,
        CdkVirtualForOf,
        CdkVirtualScrollViewport,
        LetDirective,
        UserFormGenericComponent,
        FilterFusePipe,
        FormErrorMessagePipe,
        SelectSingleComponent
    ],
    templateUrl: './user-form-company.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserFormCompanyComponent implements UserFormComponent, OnInit {
    @Input() singleColumn = false;

    @ViewChild(UserFormGenericComponent, { static: true })
    childFormComponent!: UserFormGenericComponent;

    companyIdForm!: FormControl<number | null>;
    searchCompanyForm!: FormControl<string | null>;
    companies$ = this.userFormStore.companies$;

    get form(): FormGroup<UserFormCompanyFormGroup> {
        return this.fb.group({
            ...this.childFormComponent.form.controls,
            companyId: this.companyIdForm
        });
    }

    get formValue(): UserFormCompanyFormModel {
        return {
            ...this.childFormComponent.formValue,
            companyId: this.companyIdForm.value
        };
    }

    constructor(
        private readonly fb: FormBuilder,
        private readonly userFormStore: UserFormStore
    ) {}

    ngOnInit(): void {
        this.loadData();
        this.initForm();
    }

    private loadData(): void {
        this.userFormStore.loadCompanies();
    }

    private initForm(): void {
        this.companyIdForm = this.fb.control<number | null>(null, [Validators.required]);
        this.searchCompanyForm = this.fb.control<string | null>(null, [Validators.minLength(3)]);
    }
}
