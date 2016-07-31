import { Injectable } from '@angular/core';
import { Http, Response, URLSearchParams } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';

import { MeanTweet } from './mean-tweet';

@Injectable()
export class MeanTweetService {
  private meanTweetsUrl = '/api/get_tweets';

  constructor(private http: Http) {
  }

  getMeanTweetsForCelebrity(celebrityName: string): Observable<MeanTweet[]> {
    let params = new URLSearchParams();
    params.set('celebrity', celebrityName);
    return this.http.get(this.meanTweetsUrl, { search: params })
                    .map(this.extractData)
                    .catch(this.handleError);
  }

  private extractData(res: Response): MeanTweet[] {
    let body = res.json();
    if (!body) {
      return [];
    }
    return body.map(meanTweet => ({
      id: meanTweet.Id,
      score: meanTweet.Score
    }));
  }

  private handleError(error: any) {
    let errMsg = (error.message) ? error.message :
      error.status ? `${error.status} - ${error.statusText}` : 'Server error';
    console.error(errMsg);
    return Observable.throw(errMsg);
  }
}
