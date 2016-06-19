import { describe, it, expect, async, inject, beforeEachProviders } from '@angular/core/testing';
import { TestComponentBuilder } from '@angular/compiler/testing';
import { Location } from '@angular/common';
import { SpyLocation } from '@angular/common/testing';
import { Router, RouteRegistry, ROUTER_PRIMARY_COMPONENT,
    RootRouter } from '@angular/router-deprecated';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/of';


import { AppComponent } from './app.component';
import { CelebrityService } from './celebrities/celebrity.service';
import { Celebrity } from './celebrities/celebrity';

class MockCelebrityService extends CelebrityService {
  getCelebrities(): Observable<Celebrity[]> {
    return Observable.of([]);
  }
}

describe('Component: App', () => {
  // Setup
  beforeEachProviders(() => [
    RouteRegistry,
    { provide: Location, useClass: SpyLocation },
    { provide: ROUTER_PRIMARY_COMPONENT, useValue: AppComponent },
    { provide: Router, useClass: RootRouter },
    { provide: CelebrityService, useClass: MockCelebrityService }
  ]);

  it('placholder', async(inject([TestComponentBuilder], (tcb: TestComponentBuilder) => {
    tcb.createAsync(AppComponent).then(fixture => {
      // let component = fixture.componentInstance;
      // let element = fixture.nativeElement;
      expect(true).toBeTruthy();
    });
  })));
});
