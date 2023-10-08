import { UserKindLabelPipe } from './user-kind-label.pipe';

describe('UserKindLabelPipe', () => {
    it('create an instance', () => {
        const pipe = new UserKindLabelPipe();
        expect(pipe).toBeTruthy();
    });
});
