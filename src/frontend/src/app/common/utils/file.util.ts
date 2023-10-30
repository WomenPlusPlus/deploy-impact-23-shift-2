import { LocalFile } from '@app/common/models/files.model';

export function fileUrl(file: LocalFile | File | null, defaultUrl: string | null = null): string | null {
    if (!file) {
        return defaultUrl;
    }
    if (file instanceof File) {
        return URL.createObjectURL(file);
    }
    return file.url;
}
