import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TouristViewTourComponent } from './tourist-view-tour.component';

describe('TouristViewTourComponent', () => {
  let component: TouristViewTourComponent;
  let fixture: ComponentFixture<TouristViewTourComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TouristViewTourComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TouristViewTourComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
