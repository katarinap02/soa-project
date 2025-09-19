import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KeyPointsComponent } from './key-points.component';

describe('KeyPointsComponent', () => {
  let component: KeyPointsComponent;
  let fixture: ComponentFixture<KeyPointsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ KeyPointsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(KeyPointsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
