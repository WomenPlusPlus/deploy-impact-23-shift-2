import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { TemporaryComponent } from './components/temporary/temporary.component';

const routes: Routes = [
    {
        path: '',
        component: TemporaryComponent
    }
];

@NgModule({
    imports: [RouterModule.forChild(routes)]
})
export class TemporaryRoutingModule {}
