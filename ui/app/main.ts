import { bootstrap } from 'angular2/platform/browser';
import { ENV_PROVIDERS } from './platform/environment';

import { AppComponent } from './app.component';

export function main(initialHmrState?: any): Promise<any> {
  return bootstrap(AppComponent).catch(err => console.error(err));
}

if ('development' === ENV && HMR === true) {
  // activate hot module reload
  let ngHmr = require('angular2-hmr');
  ngHmr.hotModuleReplacement(main, module);
} else {
  // bootstrap when documetn is ready
  document.addEventListener('DOMContentLoaded', () => main());
}