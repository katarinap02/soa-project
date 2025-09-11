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




const routes: Routes = [
  { path: 'create-blogpost', component: CreateBlogpostComponent },
  { path: 'create-comment', component: CreateCommentComponent },
  

  // opcionalno: redirect sa početne strane
  { path: 'register' , component: RegisterComponent},
  { path: 'view-users', component: ViewUsersComponent},
  { path: '', component: LoginComponent },
  { path: 'home', component: HomeComponent,
    children: [
      { path: 'tours-overview', component: ToursOverviewComponent },
      { path: 'my-tours', component: MyToursComponent },
      { path: 'create-tour', component: CreateTourComponent },
      { path: 'tour-keypoints/:id', component: KeyPointsComponent },
     
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
