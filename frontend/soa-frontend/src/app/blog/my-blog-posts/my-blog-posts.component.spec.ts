import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyBlogPostsComponent } from './my-blog-posts.component';

describe('MyBlogPostsComponent', () => {
  let component: MyBlogPostsComponent;
  let fixture: ComponentFixture<MyBlogPostsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MyBlogPostsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MyBlogPostsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
