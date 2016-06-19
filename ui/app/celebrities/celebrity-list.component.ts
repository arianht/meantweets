import { Component, OnInit, OnDestroy } from '@angular/core';
import { MD_CARD_DIRECTIVES } from '@angular2-material/card';
import { MD_GRID_LIST_DIRECTIVES } from '@angular2-material/grid-list';

import { Celebrity } from './celebrity';
import { CelebrityService } from './celebrity.service';

@Component({
  directives: [
    MD_GRID_LIST_DIRECTIVES,
    MD_CARD_DIRECTIVES
  ],
  template: `
    <h2>Celebrities</h2>
    <md-grid-list cols="4" gutterSize="4px">
      <md-grid-tile *ngFor="let celebrity of celebrities">
        <md-card>
          <md-card-title>{{celebrity.name}}</md-card-title>
          <img md-card-image src="{{celebrity.photoUrl}}" />
        </md-card>
      </md-grid-tile>
    </md-grid-list>
  `
})
export class CelebrityListComponent implements OnInit, OnDestroy {
  celebrities: Celebrity[] = [];
  private celebritiesUnsub;

  constructor(
    private celebrityService: CelebrityService
  ) {

  }

  ngOnInit() {
    this.celebritiesUnsub = this.celebrityService.getCelebrities().subscribe(
      celebrities => this.celebrities = celebrities,
      err => console.error(err)
    );
  }

  ngOnDestroy() {
    this.celebritiesUnsub();
  }
}
