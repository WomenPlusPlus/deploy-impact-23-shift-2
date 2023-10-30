import { debounceTime, Observable, Subscription } from 'rxjs';

import { CdkFixedSizeVirtualScroll, CdkVirtualForOf, CdkVirtualScrollViewport } from '@angular/cdk/scrolling';
import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, EventEmitter, Input, OnDestroy, OnInit, Output } from '@angular/core';
import { FormBuilder, FormControl, ReactiveFormsModule, Validators } from '@angular/forms';

import { LetDirective } from '@app/common/directives/let/let.directive';
import { FilterFusePipe } from '@app/common/pipes/filter-fuse/filter-fuse.pipe';

type Control<T, K> = K extends keyof T ? T[K] : T;

@Component({
    selector: 'app-select-single',
    standalone: true,
    imports: [
        CommonModule,
        ReactiveFormsModule,
        CdkFixedSizeVirtualScroll,
        CdkVirtualForOf,
        CdkVirtualScrollViewport,
        FilterFusePipe,
        LetDirective
    ],
    styles: [
        `
            :host {
                display: contents;
            }
        `
    ],
    templateUrl: './select-single.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SelectSingleComponent<T> implements OnInit, OnDestroy {
    @Input() id!: string;
    @Input() options!: T[] | null;
    @Input() control!: FormControl<Control<T, typeof this.bindValue> | null>;
    @Input() minSearchLen = 3;
    @Input() bindLabel?: string;
    @Input() bindValue?: string;
    @Input() inputSize = 20;
    @Input() searchKeys: string[] = [];
    @Input() clearOnSelect = false;
    @Output() valueChange = new EventEmitter<Control<T, typeof this.bindValue> | null>();

    searchControl!: FormControl<string | null>;
    searchValue$!: Observable<string | null>;

    private readonly subscriptions: Subscription[] = [];

    constructor(private readonly fb: FormBuilder) {}

    ngOnInit(): void {
        this.initForm();
        this.initSubscriptions();
    }

    ngOnDestroy(): void {
        this.subscriptions.forEach((s) => s.unsubscribe());
    }

    onSelect(value: T): void {
        this.control.markAsTouched();
        this.control.setValue(this.bindValue ? value[this.bindValue as keyof T] : (value as any));

        if (this.clearOnSelect) {
            this.searchControl.setValue(null);
        } else {
            this.searchControl.setValue(this.bindLabel ? value[this.bindLabel as keyof T] : (value as any));
        }
    }

    onDeselect(): void {
        if (this.control.value) {
            this.control.markAsTouched();
            this.control.setValue(null);
            this.searchControl.setValue(null);
        }
    }

    private initForm(): void {
        this.searchControl = this.fb.control<string | null>(null, [Validators.minLength(this.minSearchLen)]);
    }

    private initSubscriptions(): void {
        this.searchValue$ = this.searchControl.valueChanges.pipe(debounceTime(300));
        this.subscriptions.push(
            this.control.valueChanges.subscribe((value) => {
                this.valueChange.emit(value);
                this.searchControl.setValue(null);
            })
        );
    }
}
