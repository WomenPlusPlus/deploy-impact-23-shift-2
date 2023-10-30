import { IconProp } from '@fortawesome/fontawesome-svg-core';

export interface MenuItem {
    route: string[];
    label: string;
    icon: IconProp;
}
