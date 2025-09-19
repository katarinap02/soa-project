import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router'; 

import { UserView } from '../model/UserView.model';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {


   constructor(private router: Router) {}

   userRole = '';

    user: UserView | null = null;


   ngOnInit(): void {
  
    const token = localStorage.getItem('token');
    console.log(token)
    
    const userStr = localStorage.getItem('user');
    if (!token || !userStr) {
      alert('No user logged in');
      return;
    }
    console.log(userStr);
    const user = JSON.parse(userStr);
    this.user = user;

    this.userRole = this.user?.role || '';
  }


  isAdmin(): boolean {
    return this.user?.role === 'Admin';

  }

    isGuide(): boolean {
    return this.user?.role === 'Guide';

  }
 logout(): void {
    localStorage.clear();           
    this.router.navigate(['']);     
  }

  goToTours(): void {
    this.router.navigate(['home/tours-overview']);
  }

 goToMyTours(): void {
     this.router.navigate(['home/my-tours']);
   }
  
  goToPositionSimulator(): void {
    this.router.navigate(['home/position-simulator']);
  }

}
