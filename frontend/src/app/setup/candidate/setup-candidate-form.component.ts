import { CommonModule } from '@angular/common';
import { AfterViewInit, ChangeDetectionStrategy, Component, Input, ViewChild } from '@angular/core';
import { FormGroup, FormsModule } from '@angular/forms';

import { UserFormCandidateComponent } from '@app/admin/users/form/candidate/user-form-candidate.component';
import {
    UserFormCandidateFormGroup,
    UserFormComponent,
    UserFormGroup,
    UserFormModel
} from '@app/admin/users/form/common/models/user-form.model';
import { ProfileSetup } from '@app/common/models/profile.model';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { SetupCandidateFormStore } from '@app/setup/candidate/setup-candidate-form.store';

@Component({
    selector: 'app-setup-candidate-form',
    standalone: true,
    imports: [CommonModule, UserFormCandidateComponent, UserKindLabelPipe, FormsModule],
    providers: [SetupCandidateFormStore],
    templateUrl: './setup-candidate-form.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SetupCandidateFormComponent implements AfterViewInit {
    @Input() profile!: ProfileSetup;

    @ViewChild('childFormEl', { static: false })
    childFormComponent?: UserFormComponent<UserFormGroup, UserFormModel>;

    vm$ = this.setupCandidateStore.vm$;

    get childForm(): FormGroup<UserFormGroup> | undefined {
        return this.childFormComponent?.form;
    }

    constructor(private readonly setupCandidateStore: SetupCandidateFormStore) {}

    ngAfterViewInit(): void {
        this.initFormValue();
    }

    onSubmit(): void {
        if (!this.childFormComponent?.form.valid) {
            return;
        }
        this.setupCandidateStore.submitForm({
            ...this.childFormComponent.formValue,
            kind: this.profile.kind
        });
    }

    private initFormValue(): void {
        const form = this.childForm as FormGroup<UserFormCandidateFormGroup> | undefined;
        if (!form || !this.profile) {
            return;
        }
        form.patchValue({ details: { email: this.profile.email } });
    }
}
