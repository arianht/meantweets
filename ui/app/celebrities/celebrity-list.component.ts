import { Component, OnInit } from '@angular/core';
import { ROUTER_DIRECTIVES } from '@angular/router';
import { MD_CARD_DIRECTIVES } from '@angular2-material/card';
import { MD_GRID_LIST_DIRECTIVES } from '@angular2-material/grid-list';

import { Observable } from 'rxjs/Observable';

import { Celebrity } from './celebrity';
import { CelebrityService } from './celebrity.service';

@Component({
  directives: [
    MD_GRID_LIST_DIRECTIVES,
    MD_CARD_DIRECTIVES,
    ROUTER_DIRECTIVES
  ],
  styles: [
    `
      .md-card {
        cursor: pointer;
      }
    `
  ],
  template: `
    <h2>Celebrities</h2>
    <md-grid-list cols="4" gutterSize="4px">
      <md-grid-tile *ngFor="let celebrity of celebrities | async">
        <md-card routerLink="/mean-tweets/{{celebrity.name.split(' ').join('+')}}">
          <md-card-title>{{celebrity.name}}</md-card-title>
          <img md-card-image src="{{celebrity.photoUrl}}" />
        </md-card>
      </md-grid-tile>
    </md-grid-list>
  `
})
export class CelebrityListComponent implements OnInit {
  celebrities: Observable<Celebrity[]>;

  constructor(
    private celebrityService: CelebrityService
  ) {
  }

  ngOnInit() {
    this.celebrities = this.celebrityService.getCelebrities();
  }
}
