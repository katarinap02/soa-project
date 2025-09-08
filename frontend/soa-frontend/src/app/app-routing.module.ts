import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateBlogpostComponent } from './blog/create-blogpost/create-blogpost.component';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';
import { CreateTourComponent } from './tour/create-tour/create-tour.component';
import { ToursOverviewComponent } from './tour/tour-overview/tour-overview.component';



const routes: Routes = [
  { path: 'create-blogpost', component: CreateBlogpostComponent },
  { path: 'create-comment', component: CreateCommentComponent },
  { path: 'create-tour', component: CreateTourComponent},
  { path: 'tours-overview', component: ToursOverviewComponent}

  // opcionalno: redirect sa početne strane
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
