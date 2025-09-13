import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewTourMapComponent } from './view-tour-map.component';

describe('ViewTourMapComponent', () => {
  let component: ViewTourMapComponent;
  let fixture: ComponentFixture<ViewTourMapComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewTourMapComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewTourMapComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
