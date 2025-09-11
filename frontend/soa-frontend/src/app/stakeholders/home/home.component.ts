import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router'; 

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

   constructor(private router: Router) {}

   userRole = '';

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

    this.userRole = user.role;

    alert(`Token: ${token}\nUsername: ${user.username}\nRole: ${user.role}`);
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
}
