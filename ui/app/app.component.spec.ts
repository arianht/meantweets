import { describe, it, expect, async, inject, beforeEachProviders,
  TestComponentBuilder } from 'angular2/testing';
import { AppComponent } from './app.component';

describe('Component: App', () => {
  // Setup
  beforeEachProviders(() => [
  ]);

  it('placholder', async(inject([TestComponentBuilder], (tcb: TestComponentBuilder) => {
    tcb.createAsync(AppComponent).then(fixture => {
      let component = fixture.componentInstance;
      let element = fixture.nativeElement;
      expect(true).toBeTruthy();
    });
  })));
})
