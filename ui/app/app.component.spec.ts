import { describe, it, expect, async, inject, beforeEachProviders } from '@angular/core/testing';
  import { TestComponentBuilder } from '@angular/compiler/testing';
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
});
