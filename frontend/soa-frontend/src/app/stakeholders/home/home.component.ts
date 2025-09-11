import { Component, OnInit } from '@angular/core';
import { UserView } from '../model/UserView.model';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
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

    
  }

  isAdmin(): boolean {
    return this.user?.role === 'Administrator';
  }

}
