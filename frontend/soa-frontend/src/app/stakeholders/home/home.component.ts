import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
   ngOnInit(): void {
  
    const token = localStorage.getItem('token');

    
    const userStr = localStorage.getItem('user');
    if (!token || !userStr) {
      alert('No user logged in');
      return;
    }

    const user = JSON.parse(userStr);

    alert(`Token: ${token}\nUsername: ${user.username}\nRole: ${user.role}`);
  }

}
