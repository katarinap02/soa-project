import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewBlogPostsComponent } from './view-blog-posts.component';

describe('ViewBlogPostsComponent', () => {
  let component: ViewBlogPostsComponent;
  let fixture: ComponentFixture<ViewBlogPostsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewBlogPostsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewBlogPostsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
