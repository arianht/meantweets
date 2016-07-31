// Angular 2
// rc2 workaround
import { enableProdMode  } from '@angular/core';

// Environment Providers
let PROVIDERS = [
  // common env directives
];

if ('production' === ENV) {
  // Production
  enableProdMode();

  PROVIDERS = [
    ...PROVIDERS
    // custom providers in production
  ];
} else {
  // Development
  PROVIDERS = [
    ...PROVIDERS
    // custom providers in development
  ];

}

export const ENV_PROVIDERS = [
  ...PROVIDERS
];
