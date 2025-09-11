import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { UserView } from 'src/app/stakeholders/model/UserView.model';

@Component({
  selector: 'app-follow',
  templateUrl: './follow.component.html',
  styleUrls: ['./follow.component.css'],
})
export class FollowComponent implements OnInit {
  user: UserView | null = null;

  ngOnInit(): void {
    const userStr = localStorage.getItem('user');
    if (userStr) {
      this.user = JSON.parse(userStr) as UserView;
      console.log(this.user)
    }
  }

 }