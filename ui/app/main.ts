import { bootstrap } from '@angular/platform-browser-dynamic';
import { HTTP_PROVIDERS } from '@angular/http';

import { ENV_PROVIDERS } from './platform/environment';
import { appRouterProviders } from './app.routes';

import { AppComponent } from './app.component';

export function main(initialHmrState?: any): Promise<any> {
  return bootstrap(AppComponent, [
      ...ENV_PROVIDERS,
      appRouterProviders,
      HTTP_PROVIDERS
  ]).catch(err => console.error(err));
}

if ('development' === ENV && HMR === true) {
  // activate hot module reload
  let ngHmr = require('angular2-hmr');
  ngHmr.hotModuleReplacement(main, module);
} else {
  // bootstrap when documetn is ready
  document.addEventListener('DOMContentLoaded', () => main());
}
