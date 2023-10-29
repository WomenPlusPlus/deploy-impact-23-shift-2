import { faBuilding, faChartLine, faEnvelope, faList, faSitemap, faUser } from '@fortawesome/free-solid-svg-icons';

import { UserKindEnum } from '@app/common/models/users.model';
import { MenuItem } from '@app/core/menu/common/models/menu.model';

export const ADMIN_MENU: MenuItem[] = [
    {
        route: ['/dashboard'],
        label: 'Dashboard',
        icon: faChartLine
    },
    {
        route: ['/invitations'],
        label: 'Invitations',
        icon: faEnvelope
    },
    {
        route: ['/associations'],
        label: 'Associations',
        icon: faSitemap
    },
    {
        route: ['/companies'],
        label: 'Companies',
        icon: faBuilding
    },
    {
        route: ['/jobs'],
        label: 'Jobs Listing',
        icon: faList
    },
    {
        route: ['/users'],
        label: 'Users',
        icon: faUser
    }
];

export const CANDIDATE_MENU: MenuItem[] = [
    {
        route: ['/dashboard'],
        label: 'Dashboard',
        icon: faChartLine
    },
    {
        route: ['/associations'],
        label: 'Associations',
        icon: faSitemap
    },
    {
        route: ['/companies'],
        label: 'Companies',
        icon: faBuilding
    },
    {
        route: ['/jobs'],
        label: 'Jobs Listing',
        icon: faList
    },
    {
        route: ['/users'],
        label: 'Users',
        icon: faUser
    }
];

export const COMPANY_MENU: MenuItem[] = [
    {
        route: ['/dashboard'],
        label: 'Dashboard',
        icon: faChartLine
    },
    {
        route: ['/associations'],
        label: 'Associations',
        icon: faSitemap
    },
    {
        route: ['/companies'],
        label: 'Companies',
        icon: faBuilding
    },
    {
        route: ['/jobs'],
        label: 'Jobs Listing',
        icon: faList
    },
    {
        route: ['/users'],
        label: 'Users',
        icon: faUser
    }
];

export const ASSOCIATION_MENU: MenuItem[] = [
    {
        route: ['/dashboard'],
        label: 'Dashboard',
        icon: faChartLine
    },
    {
        route: ['/invitations'],
        label: 'Invitations',
        icon: faEnvelope
    },
    {
        route: ['/associations'],
        label: 'Associations',
        icon: faSitemap
    },
    {
        route: ['/companies'],
        label: 'Companies',
        icon: faBuilding
    },
    {
        route: ['/jobs'],
        label: 'Jobs Listing',
        icon: faList
    },
    {
        route: ['/users'],
        label: 'Users',
        icon: faUser
    }
];

export const MENU_ITEMS_BY_KIND: Record<UserKindEnum, MenuItem[]> = {
    [UserKindEnum.ADMIN]: ADMIN_MENU,
    [UserKindEnum.ASSOCIATION]: ASSOCIATION_MENU,
    [UserKindEnum.COMPANY]: COMPANY_MENU,
    [UserKindEnum.CANDIDATE]: CANDIDATE_MENU
};
