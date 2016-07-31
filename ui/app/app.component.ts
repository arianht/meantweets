import { Component } from '@angular/core';
import { ROUTER_DIRECTIVES } from '@angular/router';
import { MD_BUTTON_DIRECTIVES } from '@angular2-material/button';
import { MD_TOOLBAR_DIRECTIVES } from '@angular2-material/toolbar';

import { CelebrityService } from './celebrities/celebrity.service';
import { MeanTweetService } from './tweets/mean-tweet.service';

@Component({
  selector: 'my-app',
  directives: [
    ROUTER_DIRECTIVES,
    MD_BUTTON_DIRECTIVES,
    MD_TOOLBAR_DIRECTIVES
  ],
  providers: [
    CelebrityService,
    MeanTweetService
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
export class AppComponent {
  constructor() {
  }
}
