import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateBlogpostComponent } from './blog/create-blogpost/create-blogpost.component';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';
import { CreateTourComponent } from './tour/create-tour/create-tour.component';
import { ToursOverviewComponent } from './tour/tours-overview/tour-overview.component';
import { RegisterComponent } from './stakeholders/register/register.component';
import { ViewUsersComponent } from './stakeholders/view-users/view-users.component';
import { LoginComponent } from './stakeholders/login/login.component';
import { HomeComponent } from './stakeholders/home/home.component';
import { MyToursComponent } from './tour/my-tours/my-tours.component';
import { KeyPointsComponent } from './tour/key-points/key-points.component';
import { FollowComponent } from './followers/follow/follow.component';
import { ViewBlogPostsComponent } from './blog/view-blog-posts/view-blog-posts.component';
import { MyBlogPostsComponent } from './blog/my-blog-posts/my-blog-posts.component';
import { BlogPostDetailsComponent } from './blog/blog-post-details/blog-post-details.component';
import { ViewTourMapComponent } from './tour/view-tour-map/view-tour-map.component';
import { TouristViewTourComponent } from './tour/tourist-view-tour/tourist-view-tour.component';
import { PositionSimulatorComponent } from './tour/position-simulator/position-simulator.component';
import { TourExecutionComponent } from './tour/tour-execution/tour-execution.component';



const routes: Routes = [
  { path: 'create-blogpost', component: CreateBlogpostComponent },

  { path: 'create-comment', component: CreateCommentComponent },

  { path: 'create-tour', component: CreateTourComponent},
  { path: 'tours-overview', component: ToursOverviewComponent},

  // opcionalno: redirect sa početne strane
  { path: 'register' , component: RegisterComponent},
  { path: 'view-users', component: ViewUsersComponent},
  { path: '', component: LoginComponent },
  { path: 'home', component: HomeComponent},
  { path: 'home', component: HomeComponent,
    children: [
      { path: 'tours-overview', component: ToursOverviewComponent },
      { path: 'my-tours', component: MyToursComponent },
      { path: 'create-tour', component: CreateTourComponent },
      { path: 'tour-keypoints/:id', component: KeyPointsComponent },
      { path: 'view-map/:id', component: ViewTourMapComponent},
        { path: 'view-map-tourist/:id', component: TouristViewTourComponent},
        {path: 'follow', component:FollowComponent},
        { path: 'view-users', component: ViewUsersComponent},
        { path: 'position-simulator', component: PositionSimulatorComponent },
        { path: 'tour-execution/:id', component: TourExecutionComponent }
     
    ]
  },

  { path: 'create-comment/:id', component: CreateCommentComponent },
  { path: 'register' , component: RegisterComponent},
  
  { path: '', component: LoginComponent },
  {path: 'home', component: HomeComponent},
  
  { path: 'view-blogposts', component: ViewBlogPostsComponent},
   { path: 'view-my-blogposts', component: MyBlogPostsComponent},
    { path: 'post-details/:id', component: BlogPostDetailsComponent }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
