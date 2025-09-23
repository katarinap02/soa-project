import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewPostsNewComponent } from './view-posts-new.component';

describe('ViewPostsNewComponent', () => {
  let component: ViewPostsNewComponent;
  let fixture: ComponentFixture<ViewPostsNewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewPostsNewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewPostsNewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
