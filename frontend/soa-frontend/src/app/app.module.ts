import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatCardModule } from '@angular/material/card';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { RouterModule } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { CreateBlogpostComponent } from './blog/create-blogpost/create-blogpost.component';

import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';


import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';
import { CreateTourComponent } from './tour/create-tour/create-tour.component';
import { ToursOverviewComponent } from './tour/tours-overview/tour-overview.component';
import { RegisterComponent } from './stakeholders/register/register.component';
import { MatSelectModule } from '@angular/material/select';
import { ViewUsersComponent } from './stakeholders/view-users/view-users.component';
import { LoginComponent } from './stakeholders/login/login.component';
import { HomeComponent } from './stakeholders/home/home.component';

import { MyToursComponent } from './tour/my-tours/my-tours.component';
import { KeyPointsComponent } from './tour/key-points/key-points.component';
import { ReviewDialogComponent } from './tour/review-dialog/review-dialog.component';
import { ReviewComponent } from './tour/review/review.component';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatDialogModule } from '@angular/material/dialog';
import { MatCheckboxModule } from '@angular/material/checkbox';


import { FollowComponent } from './followers/follow/follow.component';
import { ViewBlogPostsComponent } from './blog/view-blog-posts/view-blog-posts.component';
import { BlogDetailsComponent } from './blog/blog-details/blog-details.component';
import { MyBlogPostsComponent } from './blog/my-blog-posts/my-blog-posts.component';
import { BlogPostDetailsComponent } from './blog/blog-post-details/blog-post-details.component';
import { ViewTourMapComponent } from './tour/view-tour-map/view-tour-map.component';
import { MatTableModule } from '@angular/material/table';

@NgModule({
  declarations: [
    AppComponent,
    CreateBlogpostComponent,
    CreateCommentComponent,
    CreateTourComponent,
    ToursOverviewComponent,
    RegisterComponent,
    ViewUsersComponent,
    LoginComponent,
    HomeComponent,
    MyToursComponent,
    KeyPointsComponent,
    ReviewDialogComponent,
    ReviewComponent,
    FollowComponent,
    ViewBlogPostsComponent,
    BlogDetailsComponent,
    MyBlogPostsComponent,
    BlogPostDetailsComponent,
    ViewTourMapComponent,
 
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    BrowserAnimationsModule,
    MatCardModule,
    MatDatepickerModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatCardModule,
    MatIconModule,
    MatDividerModule,
    MatButtonModule,
    MatSelectModule,
    BrowserAnimationsModule,
    MatDialogModule,
     MatCheckboxModule,
     MatCardModule,
     MatButtonModule,
     MatTableModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
