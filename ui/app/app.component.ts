import { Component } from '@angular/core';
import { RouteConfig, ROUTER_DIRECTIVES } from '@angular/router-deprecated';
import { MD_BUTTON_DIRECTIVES } from '@angular2-material/button';
import { MD_TOOLBAR_DIRECTIVES } from '@angular2-material/toolbar';

import { CelebrityService } from './celebrities/celebrity.service';
import { CelebrityListComponent } from './celebrities/celebrity-list.component';

@Component({
  selector: 'my-app',
  directives: [
    ROUTER_DIRECTIVES,
    MD_BUTTON_DIRECTIVES,
    MD_TOOLBAR_DIRECTIVES
  ],
  providers: [
    CelebrityService
  ],
  // styles: [
    // require('normalize.css'),
  // ],
  template: `
    <md-toolbar color="primary">
      <span>Mean Tweets</span>
    </md-toolbar>
    <router-outlet>
  `
})
@RouteConfig([
  {
    path: '/celebrty-list',
    name: 'CelebrityList',
    component: CelebrityListComponent,
    useAsDefault: true
  }
])
export class AppComponent {
  constructor() {
  }
}
