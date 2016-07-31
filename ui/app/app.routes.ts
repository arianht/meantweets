import { provideRouter, RouterConfig } from '@angular/router';

import { CelebrityListComponent } from './celebrities/celebrity-list.component';
import { MeanTweetListComponent } from './tweets/mean-tweet-list.component';

const routes: RouterConfig = [
  {
    path: '',
    redirectTo: '/celebrity-list',
    pathMatch: 'full'
  },
  {
    path: 'celebrity-list',
    component: CelebrityListComponent
  },
  {
    path: 'mean-tweets/:celebrityName',
    component: MeanTweetListComponent
  }
];

export const appRouterProviders = [
  provideRouter(routes)
];
