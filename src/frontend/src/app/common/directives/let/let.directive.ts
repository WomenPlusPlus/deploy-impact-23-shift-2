import { Directive, Input, TemplateRef, ViewContainerRef } from '@angular/core';

interface LetContext<T> {
    appLet: T;
}

@Directive({
    standalone: true,
    selector: '[appLet]'
})
export class LetDirective<T> {
    private _context: LetContext<T> = { appLet: null as T };

    constructor(_viewContainer: ViewContainerRef, _templateRef: TemplateRef<LetContext<T>>) {
        _viewContainer.createEmbeddedView(_templateRef, this._context);
    }

    @Input()
    set appLet(value: T) {
        this._context.appLet = value;
    }
}
