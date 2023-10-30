import { CdkFixedSizeVirtualScroll, CdkVirtualForOf, CdkVirtualScrollViewport } from '@angular/cdk/scrolling';
import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { FilterFusePipe } from '@app/common/pipes/filter-fuse/filter-fuse.pipe';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { SelectSingleComponent } from '@app/ui/select-single/select-single.component';
import {
    UserFormAssociationFormGroup,
    UserFormAssociationFormModel,
    UserFormComponent
} from '@app/users/form/common/models/user-form.model';
import { UserFormGenericComponent } from '@app/users/form/generic/user-form-generic.component';
import { UserFormStore } from '@app/users/form/user-form.store';

@Component({
    selector: 'app-user-form-association',
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
    templateUrl: './user-form-association.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserFormAssociationComponent implements UserFormComponent, OnInit {
    @ViewChild(UserFormGenericComponent, { static: true })
    childFormComponent!: UserFormGenericComponent;

    associationIdForm!: FormControl<number | null>;
    searchAssociationForm!: FormControl<string | null>;
    associations$ = this.userFormStore.associations$;

    get form(): FormGroup<UserFormAssociationFormGroup> {
        return this.fb.group({
            ...this.childFormComponent.form.controls,
            associationId: this.associationIdForm
        });
    }

    get formValue(): UserFormAssociationFormModel {
        return {
            ...this.childFormComponent.formValue,
            associationId: this.associationIdForm.value
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
        this.userFormStore.loadAssociations();
    }

    private initForm(): void {
        this.associationIdForm = this.fb.control<number | null>(null, [Validators.required]);
        this.searchAssociationForm = this.fb.control<string | null>(null, [Validators.minLength(3)]);
    }
}
