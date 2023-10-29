import { filter, take } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit, ViewChild } from '@angular/core';
import { FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';

import { LetDirective } from '@app/common/directives/let/let.directive';
import {
    AssociationUserDetails,
    CandidateDetails,
    CompanyUserDetails,
    UserDetails,
    UserKindEnum
} from '@app/common/models/users.model';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';
import { UserKindLabelPipe } from '@app/common/pipes/user-kind-label/user-kind-label.pipe';
import { toInputDateValue } from '@app/common/utils/date.util';
import { ContentErrorComponent } from '@app/ui/content-error/content-error.component';
import { ContentLoadingComponent } from '@app/ui/content-loading/content-loading.component';
import { UserFormAssociationComponent } from '@app/users/form/association/user-form-association.component';
import { UserFormCandidateComponent } from '@app/users/form/candidate/user-form-candidate.component';
import {
    UserFormAssociationFormGroup,
    UserFormCandidateFormGroup,
    UserFormCompanyFormGroup,
    UserFormComponent,
    UserFormGroup,
    UserFormModel
} from '@app/users/form/common/models/user-form.model';
import { UserFormCompanyComponent } from '@app/users/form/company/user-form-company.component';
import { EditUserStore } from '@app/users/form/edit-user.store';
import { UserFormGenericComponent } from '@app/users/form/generic/user-form-generic.component';
import { UserFormStore } from '@app/users/form/user-form.store';

@Component({
    selector: 'app-edit-user',
    standalone: true,
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        RouterModule,
        UserFormGenericComponent,
        UserFormCandidateComponent,
        FormErrorMessagePipe,
        LetDirective,
        UserKindLabelPipe,
        UserFormCompanyComponent,
        UserFormAssociationComponent,
        ContentErrorComponent,
        ContentLoadingComponent
    ],
    providers: [UserFormStore, EditUserStore],
    templateUrl: './edit-user.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class EditUserComponent implements OnInit {
    @ViewChild('childFormEl', { static: false })
    childFormComponent?: UserFormComponent<UserFormGroup, UserFormModel>;

    vm$ = this.editUserStore.vm$;

    get childForm(): FormGroup<UserFormGroup> | undefined {
        return this.childFormComponent?.form;
    }

    get userEmail(): string {
        return this.childFormComponent?.form.controls.details.controls.email.value || 'N/A';
    }

    protected readonly userKindsEnum = UserKindEnum;

    constructor(
        private readonly router: Router,
        private readonly route: ActivatedRoute,
        private readonly editUserStore: EditUserStore
    ) {}

    ngOnInit(): void {
        this.loadData();
    }

    onSubmit(user: UserDetails): void {
        if (!this.childFormComponent?.form.valid) {
            return;
        }
        this.editUserStore.submitForm({
            ...this.childFormComponent.formValue,
            kind: user.kind,
            id: user.id
        });
    }

    private loadData(): void {
        const id = this.route.snapshot.paramMap.get('id');
        if (!id || isNaN(+id)) {
            this.router.navigate(['..']);
            return;
        }
        this.editUserStore
            .getUser(+id)
            .add(() =>
                this.editUserStore.user$
                    .pipe(filter(Boolean), take(1))
                    .subscribe((user) => setTimeout(() => this.setUserValue(user)))
            );
    }

    private setUserValue(user: UserDetails): void {
        switch (user.kind) {
            case UserKindEnum.ADMIN:
                return this.setAdminUserValue(user);
            case UserKindEnum.ASSOCIATION:
                return this.setAssociationUserValue(user as AssociationUserDetails);
            case UserKindEnum.CANDIDATE:
                return this.setCandidateUserValue(user as CandidateDetails);
            case UserKindEnum.COMPANY:
                return this.setCompanyUserValue(user as CompanyUserDetails);
        }
    }

    private setAdminUserValue(user: UserDetails): void {
        const form = this.childForm;
        if (!form) {
            return;
        }
        form.setValue({
            details: {
                firstName: user.firstName,
                lastName: user.lastName,
                preferredName: user.preferredName,
                email: user.email,
                phoneNumber: user.phoneNumber,
                birthDate: user.birthDate ? toInputDateValue(user.birthDate) : null,
                photo: user.photo || null
            },
            social: {
                linkedInUrl: user.linkedInUrl,
                githubUrl: user.githubUrl,
                portfolioUrl: user.portfolioUrl
            }
        });
    }

    private setAssociationUserValue(user: AssociationUserDetails): void {
        const form = this.childForm as FormGroup<UserFormAssociationFormGroup> | undefined;
        if (!form) {
            return;
        }
        form.setValue({
            associationId: user.associationId,
            details: {
                firstName: user.firstName,
                lastName: user.lastName,
                preferredName: user.preferredName,
                email: user.email,
                phoneNumber: user.phoneNumber,
                birthDate: user.birthDate ? toInputDateValue(user.birthDate) : null,
                photo: user.photo || null
            },
            social: {
                linkedInUrl: user.linkedInUrl,
                githubUrl: user.githubUrl,
                portfolioUrl: user.portfolioUrl
            }
        });
    }

    private setCandidateUserValue(user: CandidateDetails): void {
        const form = this.childForm as FormGroup<UserFormCandidateFormGroup> | undefined;
        if (!form) {
            return;
        }
        form.setValue({
            details: {
                firstName: user.firstName,
                lastName: user.lastName,
                preferredName: user.preferredName,
                email: user.email,
                phoneNumber: user.phoneNumber,
                birthDate: user.birthDate ? toInputDateValue(user.birthDate) : null,
                photo: user.photo || null
            },
            social: {
                linkedInUrl: user.linkedInUrl,
                githubUrl: user.githubUrl,
                portfolioUrl: user.portfolioUrl
            },
            job: {
                yearsOfExperience: user.yearsOfExperience,
                jobStatus: user.jobStatus,
                seekJobType: user.seekJobType,
                seekCompanySize: user.seekCompanySize,
                seekLocations: user.seekLocations,
                seekLocationType: user.seekLocationType,
                seekSalary: user.seekSalary,
                seekValues: user.seekValues,
                workPermit: user.workPermit,
                noticePeriod: user.noticePeriod
            },
            technical: {
                spokenLanguages: user.spokenLanguages,
                skills: user.skills,
                cv: user.cv || null,
                attachments: user.attachments,
                video: user.video || null,
                educationHistory: user.educationHistory.map((history) => ({
                    title: history.title,
                    description: history.description,
                    entity: history.entity,
                    fromDate: new Date(history.fromDate),
                    toDate: history.toDate ? new Date(history.toDate) : null,
                    onGoing: !history.toDate
                })),
                employmentHistory: user.employmentHistory.map((history) => ({
                    title: history.title,
                    description: history.description,
                    company: history.company,
                    fromDate: new Date(history.fromDate),
                    toDate: history.toDate ? new Date(history.toDate) : null,
                    onGoing: !history.toDate
                }))
            }
        });
    }

    private setCompanyUserValue(user: CompanyUserDetails): void {
        const form = this.childForm as FormGroup<UserFormCompanyFormGroup> | undefined;
        if (!form) {
            return;
        }
        form.setValue({
            companyId: user.companyId,
            details: {
                firstName: user.firstName,
                lastName: user.lastName,
                preferredName: user.preferredName,
                email: user.email,
                phoneNumber: user.phoneNumber,
                birthDate: user.birthDate ? toInputDateValue(user.birthDate) : null,
                photo: user.photo || null
            },
            social: {
                linkedInUrl: user.linkedInUrl,
                githubUrl: user.githubUrl,
                portfolioUrl: user.portfolioUrl
            }
        });
    }
}
