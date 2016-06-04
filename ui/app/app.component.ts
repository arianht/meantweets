import { Component } from '@angular/core';
import { MD_BUTTON_DIRECTIVES } from '@angular2-material/button';

@Component({
  directives: [MD_BUTTON_DIRECTIVES],
  selector: 'my-app',
  // styles: [
    // require('normalize.css'),
  // ],
  template: `
    <h1>Hello World!</h1>
    <button md-raised-button>Test Button!</button>
  `
})
export class AppComponent {
  constructor() {
  }
}
