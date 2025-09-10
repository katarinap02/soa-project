import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { HttpClientModule } from '@angular/common/http';
import { CreateBlogpostComponent } from './blog/create-blogpost/create-blogpost.component';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';
<<<<<<< HEAD
import { RegisterComponent } from './stakeholders/register/register.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSelectModule } from '@angular/material/select';
import { ViewUsersComponent } from './stakeholders/view-users/view-users.component';

=======
import { LoginComponent } from './stakeholders/login/login.component';

import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomeComponent } from './stakeholders/home/home.component';
>>>>>>> 80e02e5f6cd0c088f7d99bd099f14a47f2cc5cbb

@NgModule({
  declarations: [
    AppComponent,
    CreateBlogpostComponent,
    CreateCommentComponent,
<<<<<<< HEAD
    RegisterComponent,
    ViewUsersComponent
=======
    LoginComponent,
    HomeComponent
>>>>>>> 80e02e5f6cd0c088f7d99bd099f14a47f2cc5cbb
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
<<<<<<< HEAD
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatCardModule,
    MatButtonModule,
    MatSelectModule,
    BrowserAnimationsModule,
    
=======
     BrowserAnimationsModule,
    FormsModule,
    MatInputModule,
    MatButtonModule,
    MatFormFieldModule,
    HttpClientModule,
    AppRoutingModule
>>>>>>> 80e02e5f6cd0c088f7d99bd099f14a47f2cc5cbb
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
