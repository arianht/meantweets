import { Component, OnInit, OnDestroy, ElementRef } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { MD_CARD_DIRECTIVES } from '@angular2-material/card';
import { MD_GRID_LIST_DIRECTIVES } from '@angular2-material/grid-list';

import { MeanTweetService } from './mean-tweet.service';

@Component({
  directives: [
    MD_GRID_LIST_DIRECTIVES,
    MD_CARD_DIRECTIVES
  ],
  template: `
    <h2>Mean Tweets for {{celebrityName.split('+').join(' ')}}</h2>
  `
})
export class MeanTweetListComponent implements OnInit, OnDestroy {
  celebrityName: string;
  private meanTweetsUnsub;

  constructor(
    private route: ActivatedRoute,
    private meanTweetService: MeanTweetService,
    private elementRef: ElementRef
  ) {
  }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.celebrityName = (<any> params).celebrityName;
      this.meanTweetsUnsub = this.meanTweetService.getMeanTweetsForCelebrity(this.celebrityName)
          .subscribe(
            tweets => {
              tweets.forEach(tweet => {
                twttr.widgets.createTweet(tweet.id, this.elementRef.nativeElement, {}).then();
              });
            },
            err => console.error(err)
          );
    });
  }

  ngOnDestroy() {
    this.meanTweetsUnsub.unsubscribe();
  }
}
