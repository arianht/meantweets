// Angular 2
import { enableProdMode } from '@angular/core';

// Environment Providers
let PROVIDERS = [];

if ('production' === ENV) {
  // Production
  enableProdMode();
} else {
  // Development
}

export const ENV_PROVIDERS = [
  ...PROVIDERS
];
