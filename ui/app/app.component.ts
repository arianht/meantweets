import { Component } from '@angular/core';
import { MD_BUTTON_DIRECTIVES } from '@angular2-material/button';
import { MD_TOOLBAR_DIRECTIVES } from '@angular2-material/toolbar';

@Component({
  directives: [
    MD_BUTTON_DIRECTIVES,
    MD_TOOLBAR_DIRECTIVES
  ],
  selector: 'my-app',
  // styles: [
    // require('normalize.css'),
  // ],
  template: `
    <md-toolbar color="primary">
      <span>Mean Tweets</span>
    </md-toolbar>
    <h1>Hello World!</h1>
    <button md-raised-button>Test Button!</button>
  `
})
export class AppComponent {
  constructor() {
  }
}
