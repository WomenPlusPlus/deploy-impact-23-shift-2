import { Observable, share } from 'rxjs';

import { HttpClient } from '@angular/common/http';
import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';

interface CardData {
    title: string;
    content: string;
    imageUrl: string;
}

@Component({
    selector: 'app-card',
    templateUrl: './card.component.html',
    styleUrls: ['./card.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CardComponent implements OnInit {
    data$!: Observable<CardData>;

    constructor(private readonly httpClient: HttpClient) {}

    ngOnInit(): void {
        this.data$ = this.httpClient.get<CardData>('https://shift2-deployimpact.xyz/').pipe(share());
    }
}
