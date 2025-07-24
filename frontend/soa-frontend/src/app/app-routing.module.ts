import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateBlogpostComponent } from './blog/create-blogpost/create-blogpost.component';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';



const routes: Routes = [
  { path: 'create-blogpost', component: CreateBlogpostComponent },
  { path: 'create-comment', component: CreateCommentComponent }
  // opcionalno: redirect sa početne strane
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
