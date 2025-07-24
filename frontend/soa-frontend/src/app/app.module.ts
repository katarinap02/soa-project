import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { HttpClientModule } from '@angular/common/http';
import { CreateBlogpostComponent } from './blog/create-blogpost/create-blogpost.component';

import { ReactiveFormsModule } from '@angular/forms';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';
@NgModule({
  declarations: [
    AppComponent,
    CreateBlogpostComponent,
    CreateCommentComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
