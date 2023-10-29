import { BehaviorSubject, map, merge, Observable } from 'rxjs';

import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input, OnInit } from '@angular/core';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';

import { CreateAssociationFormGroup } from '@app/admin/associations/create-association/common/models/create-association.model';
import { LetDirective } from '@app/common/directives/let/let.directive';
import { FormErrorMessagePipe } from '@app/common/pipes/form-error-message/form-error-message.pipe';

const DEFAULT_PHOTO_URL = 'assets/profile-picture-default-form.png';

@Component({
    selector: 'app-association-form',
    standalone: true,
    imports: [CommonModule, ReactiveFormsModule, FormErrorMessagePipe, LetDirective],
    templateUrl: './association-form.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class AssociationFormComponent implements OnInit {
    @Input() form!: FormGroup<CreateAssociationFormGroup>;

    @Input() set previousImage(value: string) {
        value && this.previousImage$.next(value);
    }

    imageURL$!: Observable<string>;

    private previousImage$ = new BehaviorSubject<string>(DEFAULT_PHOTO_URL);

    ngOnInit(): void {
        this.initSubscriptions();
    }

    onFileSelected(event: Event): void {
        const file = (event.target as HTMLInputElement).files?.[0];
        this.form.controls.logo.setValue(file || null);
    }

    private initSubscriptions(): void {
        this.imageURL$ = merge(
            this.previousImage$,
            this.form.controls.logo.valueChanges.pipe(
                map((file: File | null) => {
                    if (!file) {
                        return DEFAULT_PHOTO_URL;
                    }
                    try {
                        return URL.createObjectURL(file);
                    } catch (error) {
                        return DEFAULT_PHOTO_URL;
                    }
                })
            )
        );
    }
}
