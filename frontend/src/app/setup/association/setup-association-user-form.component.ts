import { CommonModule } from '@angular/common';
import { AfterViewInit, ChangeDetectionStrategy, Component, Input, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, Validators } from '@angular/forms';

import { CreateAssociationFormGroup } from '@app/admin/associations/create-association/common/models/create-association.model';
import { AssociationFormComponent } from '@app/admin/associations/form/association-form.component';
import {
    UserFormAssociationFormGroup,
    UserFormComponent,
    UserFormGroup,
    UserFormModel
} from '@app/admin/users/form/common/models/user-form.model';
import { UserFormGenericComponent } from '@app/admin/users/form/generic/user-form-generic.component';
import { UserFormStore } from '@app/admin/users/form/user-form.store';
import { ProfileSetup } from '@app/common/models/profile.model';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { SetupAssociationUserFormStore } from '@app/setup/association/setup-association-user-form.store';

@Component({
    selector: 'app-setup-association-user-form',
    standalone: true,
    imports: [CommonModule, UserKindLabelPipe, FormsModule, AssociationFormComponent, UserFormGenericComponent],
    providers: [SetupAssociationUserFormStore, UserFormStore],
    templateUrl: './setup-association-user-form.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SetupAssociationUserFormComponent implements OnInit, AfterViewInit {
    @Input() profile: ProfileSetup = {
        email: 'test@t.c',
        association: {
            id: 2,
            name: 'Test Association',
            imageUrl: {
                name: 'test',
                url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/SMPTE_Color_Bars.svg/1200px-SMPTE_Color_Bars.svg.png'
            },
            websiteUrl: 'http://test-association-link',
            focus: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. '
        }
    } as ProfileSetup;

    @ViewChild('childFormEl', { static: false })
    childFormComponent?: UserFormComponent<UserFormGroup, UserFormModel>;

    vm$ = this.setupAssociationUserStore.vm$;
    associationForm!: FormGroup<CreateAssociationFormGroup>;

    get childForm(): FormGroup<UserFormGroup> | undefined {
        return this.childFormComponent?.form;
    }

    constructor(
        private readonly fb: FormBuilder,
        private readonly setupAssociationUserStore: SetupAssociationUserFormStore
    ) {}

    ngOnInit(): void {
        this.initForm();
    }

    ngAfterViewInit(): void {
        this.initFormValue();
    }

    onSubmit(): void {
        if (!this.childFormComponent?.form.valid || !this.associationForm.valid) {
            return;
        }
        const associationData = new FormData();
        const formValue = this.associationForm.getRawValue();
        for (const key of Object.keys(formValue)) {
            associationData.append(key, formValue[key as keyof CreateAssociationFormGroup] as string);
        }
        this.setupAssociationUserStore.submitForm({
            associationId: this.profile.association?.id,
            user: {
                ...this.childFormComponent.formValue,
                kind: this.profile.kind
            },
            association: associationData
        });
    }

    private initForm(): void {
        this.associationForm = this.fb.group({
            name: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(256)]),
            logo: this.fb.control<File | null>(null),
            websiteUrl: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(512)]),
            focus: this.fb.control<string | null>(null, [Validators.required, Validators.maxLength(1024)])
        });
    }

    private initFormValue(): void {
        const form = this.childForm as FormGroup<UserFormAssociationFormGroup> | undefined;
        if (!form || !this.profile) {
            return;
        }
        form.patchValue({ details: { email: this.profile.email } });
        form.controls.details.controls.email.disable({ emitEvent: false });

        if (!this.profile.association) {
            return;
        }
        this.associationForm.patchValue(this.profile.association);
        this.associationForm.disable({ emitEvent: false });
    }
}
