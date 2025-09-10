import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateBlogpostComponent } from './blog/create-blogpost/create-blogpost.component';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';
import { LoginComponent } from './stakeholders/login/login.component';
import { HomeComponent } from './stakeholders/home/home.component';



const routes: Routes = [
  { path: 'create-blogpost', component: CreateBlogpostComponent },
  { path: 'create-comment', component: CreateCommentComponent },
  { path: '', component: LoginComponent },
  {path: 'home', component: HomeComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
