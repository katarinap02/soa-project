import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateBlogpostComponent } from './create-blogpost.component';

describe('CreateBlogpostComponent', () => {
  let component: CreateBlogpostComponent;
  let fixture: ComponentFixture<CreateBlogpostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateBlogpostComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateBlogpostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
