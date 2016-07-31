import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { MD_CARD_DIRECTIVES } from '@angular2-material/card';
import { MD_GRID_LIST_DIRECTIVES } from '@angular2-material/grid-list';

import { Observable } from 'rxjs/Observable';

import { MeanTweet } from './mean-tweet';
import { MeanTweetService } from './mean-tweet.service';

@Component({
  directives: [
    MD_GRID_LIST_DIRECTIVES,
    MD_CARD_DIRECTIVES
  ],
  template: `
    <h2>Mean Tweets</h2>
    <ul>
      <li *ngFor="let meanTweet of meanTweets | async">
        {{meanTweet.id}}
      </li>
    </ul>
  `
})
export class MeanTweetListComponent implements OnInit {
  meanTweets: Observable<MeanTweet[]>;

  constructor(
    private route: ActivatedRoute,
    private meanTweetService: MeanTweetService
  ) {
  }

  ngOnInit() {
    this.route.params.subscribe(params => {
      let celebrityName = (<any> params).celebrityName;
      this.meanTweets = this.meanTweetService.getMeanTweetsForCelebrity(celebrityName);
    });
  }
}
