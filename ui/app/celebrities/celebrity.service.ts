import { Injectable } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/of';

import { Celebrity } from './celebrity';

import * as _ from 'lodash';

const CELEBRITIES = require('./celebrities.json');

@Injectable()
export class CelebrityService {
  getCelebrities(): Observable<Celebrity[]> {
    return Observable.of(_.map(CELEBRITIES, (photoUrl, name: string) => ({
      name,
      photoUrl
    })));
  }
}
