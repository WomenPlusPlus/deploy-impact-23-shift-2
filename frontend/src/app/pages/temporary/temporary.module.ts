import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { CardModule } from '../../modules/card/card.module';
import { TemporaryComponent } from './components/temporary/temporary.component';
import { TemporaryRoutingModule } from './temporary-routing.module';

@NgModule({
    declarations: [TemporaryComponent],
    imports: [CommonModule, TemporaryRoutingModule, CardModule],
    exports: [TemporaryComponent]
})
export class TemporaryModule {}
