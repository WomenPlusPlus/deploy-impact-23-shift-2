import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LocalFile } from '@app/common/models/files.model';
import { fileUrl } from '@app/common/utils/file.util';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faRemove } from '@fortawesome/free-solid-svg-icons';

@Component({
    selector: 'app-file-item',
    standalone: true,
    imports: [CommonModule, FontAwesomeModule],
    templateUrl: './file-item.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class FileItemComponent {
    @Input() file!: LocalFile | File;
    @Output() remove = new EventEmitter<void>();

    get name(): string {
        return this.file.name;
    }

    get url(): string | null {
        return fileUrl(this.file);
    }

    protected readonly faRemove = faRemove;
}
