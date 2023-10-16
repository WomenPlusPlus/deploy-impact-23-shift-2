import { CdkFixedSizeVirtualScroll, CdkVirtualForOf, CdkVirtualScrollViewport } from '@angular/cdk/scrolling';
import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';

import { CreateUserAssociationFormGroup } from '@app/admin/users/creation/common/models/create-user.model';
import { CreateUserStore } from '@app/admin/users/creation/create-user.store';
import { CreateUserGenericComponent } from '@app/admin/users/creation/generic/create-user-generic.component';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { Association } from '@app/common/models/associations.model';
import { FilterFusePipe } from '@app/common/pipes/filter-fuse/filter-fuse.pipe';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';

@Component({
    selector: 'app-create-user-association',
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
        FormErrorMessagePipe
    ],
    templateUrl: './create-user-association.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreateUserAssociationComponent implements OnInit {
    @ViewChild(CreateUserGenericComponent, { static: true })
    childFormComponent!: CreateUserGenericComponent;

    associationIdForm!: FormControl<number | null>;
    searchAssociationForm!: FormControl<string | null>;
    associations$ = this.createUserStore.associations$;

    get form(): FormGroup<CreateUserAssociationFormGroup> {
        return this.fb.group({
            ...this.childFormComponent.form.controls,
            associationId: this.associationIdForm
        });
    }

    constructor(
        private readonly fb: FormBuilder,
        private readonly createUserStore: CreateUserStore
    ) {}

    ngOnInit(): void {
        this.loadData();
        this.initForm();
    }

    onSelectAssociation(association: Association): void {
        const control = this.associationIdForm;
        control.markAsTouched();
        control.setValue(association.id);

        this.searchAssociationForm.setValue(association.name);
    }

    onDeselectAssociation(): void {
        const control = this.associationIdForm;
        if (control.value) {
            control.markAsTouched();
            control.setValue(null);
        }
    }

    private loadData(): void {
        this.createUserStore.loadAssociations();
    }

    private initForm(): void {
        this.associationIdForm = this.fb.control<number | null>(null, [Validators.required]);
        this.searchAssociationForm = this.fb.control<string | null>(null, [Validators.minLength(3)]);
    }
}
