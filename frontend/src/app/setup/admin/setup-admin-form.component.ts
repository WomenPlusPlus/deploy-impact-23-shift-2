import { CommonModule } from '@angular/common';
import { AfterViewInit, ChangeDetectionStrategy, Component, Input, ViewChild } from '@angular/core';
import { FormGroup, FormsModule } from '@angular/forms';

import { UserFormComponent, UserFormGroup, UserFormModel } from '@app/admin/users/form/common/models/user-form.model';
import { UserFormGenericComponent } from '@app/admin/users/form/generic/user-form-generic.component';
import { ProfileSetup } from '@app/common/models/profile.model';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { SetupAdminFormStore } from '@app/setup/admin/setup-admin-form.store';

@Component({
    selector: 'app-setup-admin-form',
    standalone: true,
    imports: [CommonModule, UserFormGenericComponent, UserKindLabelPipe, FormsModule],
    providers: [SetupAdminFormStore],
    templateUrl: './setup-admin-form.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SetupAdminFormComponent implements AfterViewInit {
    @Input() profile!: ProfileSetup;

    @ViewChild('childFormEl', { static: false })
    childFormComponent?: UserFormComponent<UserFormGroup, UserFormModel>;

    vm$ = this.setupAdminStore.vm$;

    get childForm(): FormGroup<UserFormGroup> | undefined {
        return this.childFormComponent?.form;
    }

    constructor(private readonly setupAdminStore: SetupAdminFormStore) {}

    ngAfterViewInit(): void {
        this.initFormValue();
    }

    onSubmit(): void {
        if (!this.childFormComponent?.form.valid) {
            return;
        }
        this.setupAdminStore.submitForm({
            ...this.childFormComponent.formValue,
            kind: this.profile.kind
        });
    }

    private initFormValue(): void {
        const form = this.childForm as FormGroup<UserFormGroup> | undefined;
        if (!form || !this.profile) {
            return;
        }
        form.patchValue({ details: { email: this.profile.email } });
    }
}
