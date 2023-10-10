import { FormErrorMessagePipe } from './form-error-message.pipe';

describe('FormErrorMessagePipe', () => {
    it('create an instance', () => {
        const pipe = new FormErrorMessagePipe();
        expect(pipe).toBeTruthy();
    });
});
