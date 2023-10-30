import { Pipe, PipeTransform } from '@angular/core';
import { ValidationErrors } from '@angular/forms';

@Pipe({
    name: 'formErrorMessage',
    standalone: true
})
export class FormErrorMessagePipe implements PipeTransform {
    transform(errors: ValidationErrors | null): string {
        if (!errors) {
            return 'N/A';
        }
        for (const errorKey of Object.keys(errors)) {
            switch (errorKey) {
                case 'required':
                    return 'This field is required.';
                case 'minlength':
                    return `Minimum length is ${errors[errorKey].requiredLength} characters.`;
                case 'maxlength':
                    return `Maximum length is ${errors[errorKey].requiredLength} characters.`;
                case 'pattern':
                    return 'Invalid format.';
                case 'email':
                    return 'Invalid email format.';
                case 'min':
                    return `Value must be greater than or equal to ${errors[errorKey].min}.`;
                case 'max':
                    return `Value must be less than or equal to ${errors[errorKey].max}.`;
                case 'requiredTrue':
                    return 'You must agree to the terms.';
            }
        }
        return 'Custom validation failed.';
    }
}
