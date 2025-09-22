import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { UserView } from '../model/UserView.model';
import { UserService } from '../user.service';

@Component({
  selector: 'app-view-users',
  templateUrl: './view-users.component.html',
  styleUrls: ['./view-users.component.css'],
})
export class ViewUsersComponent implements OnInit {
  users: UserView[] = [];

  constructor(private userService: UserService) {}

  ngOnInit(): void {
  this.userService.getAllUsers().subscribe({
    next: (data) => {
      console.log(data);
      this.users = data.users
        .filter((user: { role: string }) => user.role !== 'Admin')
        .map((user: any) => ({
          ...user,          
          username: user.name 
        }));
    },
    error: (err) => console.error('Error loading users', err)
  });
}


}
