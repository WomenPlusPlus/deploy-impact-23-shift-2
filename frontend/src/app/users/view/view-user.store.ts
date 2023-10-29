import { HotToastService } from '@ngneat/hot-toast';
import { exhaustMap, Observable, tap } from 'rxjs';

import { Injectable } from '@angular/core';

import { ComponentStore, tapResponse } from '@ngrx/component-store';

import { UserDetails, UserStateEnum } from '@app/common/models/users.model';
import { UsersService } from '@app/users/common/services/users.service';

export interface ViewUserState {
    user: UserDetails | null;
    loading: boolean;
    error: boolean;
    deleting: boolean;
}

const initialState: ViewUserState = {
    user: null,
    loading: false,
    deleting: false,
    error: false
};

@Injectable()
export class ViewUserStore extends ComponentStore<ViewUserState> {
    vm$ = this.select({
        user: this.select((state) => state.user),
        loading: this.select((state) => state.loading),
        error: this.select((state) => state.error),
        deleting: this.select((state) => state.deleting)
    });

    getUser = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.getListLoading()),
            exhaustMap((id: number) =>
                this.adminUsersService.getById(id).pipe(
                    tapResponse(
                        (list) => this.getListLoadedSuccess(list),
                        () => this.getListLoadedError()
                    )
                )
            )
        )
    );

    deleteItem = this.effect((trigger$: Observable<number>) =>
        trigger$.pipe(
            tap(() => this.patchState({ deleting: true })),
            exhaustMap((id: number) =>
                this.adminUsersService.deleteUser(id).pipe(
                    tapResponse(
                        () => {
                            this.patchState({
                                user: { ...this.state().user!, state: UserStateEnum.DELETED },
                                deleting: false
                            });
                        },
                        () => {
                            this.patchState({ deleting: false });
                            this.toast.error('Could not delete user! Please try again later or contact the support.');
                        }
                    )
                )
            )
        )
    );

    private getListLoading = this.updater(
        (state): ViewUserState => ({
            ...state,
            loading: true,
            error: false
        })
    );
    private getListLoadedSuccess = this.updater(
        (state, user: UserDetails): ViewUserState => ({
            ...state,
            user,
            loading: false
        })
    );
    private getListLoadedError = this.updater(
        (state): ViewUserState => ({
            ...state,
            error: true,
            loading: false
        })
    );

    constructor(
        private readonly adminUsersService: UsersService,
        private readonly toast: HotToastService
    ) {
        super(initialState);
    }
}
