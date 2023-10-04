import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';

import { CardComponent } from './components/card/card.component';

@NgModule({
    declarations: [CardComponent],
    imports: [CommonModule, HttpClientModule],
    exports: [CardComponent]
})
export class CardModule {}
