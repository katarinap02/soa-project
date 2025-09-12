import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateBlogpostComponent } from './blog/create-blogpost/create-blogpost.component';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';
import { RegisterComponent } from './stakeholders/register/register.component';
import { ViewUsersComponent } from './stakeholders/view-users/view-users.component';
import { LoginComponent } from './stakeholders/login/login.component';
import { HomeComponent } from './stakeholders/home/home.component';
import { FollowComponent } from './followers/follow/follow.component';
import { ViewBlogPostsComponent } from './blog/view-blog-posts/view-blog-posts.component';
import { MyBlogPostsComponent } from './blog/my-blog-posts/my-blog-posts.component';
import { BlogPostDetailsComponent } from './blog/blog-post-details/blog-post-details.component';




const routes: Routes = [
  { path: 'create-blogpost', component: CreateBlogpostComponent },
  { path: 'create-comment/:id', component: CreateCommentComponent },
  { path: 'register' , component: RegisterComponent},
  { path: 'view-users', component: ViewUsersComponent},
  { path: '', component: LoginComponent },
  {path: 'home', component: HomeComponent},
  {path: 'follow', component:FollowComponent},
  { path: 'view-blogposts', component: ViewBlogPostsComponent},
   { path: 'view-my-blogposts', component: MyBlogPostsComponent},
    { path: 'post-details/:id', component: BlogPostDetailsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
