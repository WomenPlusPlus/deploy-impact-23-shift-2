import { UserStateLabelPipe } from './user-state-label.pipe';

describe('UserStateLabelPipe', () => {
    it('create an instance', () => {
        const pipe = new UserStateLabelPipe();
        expect(pipe).toBeTruthy();
    });
});
