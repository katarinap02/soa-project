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
      this.users = data
        .filter((user: { role: string }) => user.role !== 'Admin')
        .map((user: any) => ({
          ...user,
          username: user.username,
          account_status: user.account_status || 'Activated',
        }));
    },
    error: (err) => console.error('Error loading users', err)
  });
}
blockUser(userUsername: string) {
  if (!confirm("Are you sure you want to block this user?")) return;

  const adminStr = localStorage.getItem('user');
  if (!adminStr) return;
  const admin = JSON.parse(adminStr);
  const adminUsername = admin.name; // or however it's stored

  const payload = {
    adminUsername: adminUsername,
    userToBlock: userUsername
  };

  console.log("Payload being sent to block user:", payload); // ✅ log it

  this.userService.blockUser(payload.adminUsername, payload.userToBlock).subscribe({
    next: res => {
      console.log("User blocked successfully", res);
      this.users = this.users.map(u =>
        u.username === userUsername ? { ...u, account_status: 'Blocked' } : u
      );
    },
    error: err => {
      console.error("Error blocking user:", err);
      alert("Failed to block user.");
    }
  });
}

}
