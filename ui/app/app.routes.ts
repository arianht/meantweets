import { provideRouter, RouterConfig } from '@angular/router';

import { CelebrityListComponent } from './celebrities/celebrity-list.component';

const routes: RouterConfig = [
  {
    path: '',
    redirectTo: '/celebrity-list',
    pathMatch: 'full'
  },
  {
    path: 'celebrity-list',
    component: CelebrityListComponent
  }
];

export const appRouterProviders = [
  provideRouter(routes)
];
