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
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CreateCommentComponent } from './blog/create-comment/create-comment.component';
import { CreateTourComponent } from './tour/create-tour/create-tour.component';
import { ToursOverviewComponent } from './tour/tour-overview/tour-overview.component';
import { RegisterComponent } from './stakeholders/register/register.component';
import { MatSelectModule } from '@angular/material/select';
import { ViewUsersComponent } from './stakeholders/view-users/view-users.component';
import { LoginComponent } from './stakeholders/login/login.component';
import { HomeComponent } from './stakeholders/home/home.component';

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
    HomeComponent
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
    
    ReactiveFormsModule,
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatCardModule,
    MatButtonModule,
    MatSelectModule,
    BrowserAnimationsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
