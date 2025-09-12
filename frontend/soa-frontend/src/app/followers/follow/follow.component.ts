import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { UserView } from 'src/app/stakeholders/model/UserView.model';
import { Profile } from '../model/profile.model';
import { FollowService } from '../follow.service';
import { UserService } from 'src/app/stakeholders/user.service';
import { RecommendationResponse } from '../model/recommendation-response.model';
import { forkJoin } from 'rxjs';

@Component({
  selector: 'app-follow',
  templateUrl: './follow.component.html',
  styleUrls: ['./follow.component.css'],
})
export class FollowComponent implements OnInit {
  userProfile: Profile | null = null;
  following: Profile[] = [];
  followers: Profile[] = [];
  recommendations: Profile[] = [];
  DEFAULT_AVATAR = 'https://static.vecteezy.com/system/resources/thumbnails/005/544/718/small_2x/profile-icon-design-free-vector.jpg';

  constructor(
    private profileService: UserService,
    private followService: FollowService
  ) {}

  ngOnInit(): void {
    const userStr = localStorage.getItem('user');
    
    if (!userStr) return;

    const user = JSON.parse(userStr);
    const userId = user.id;
    console.log(userId)

    // 1. Profil
    this.profileService.getProfile(userId).subscribe({
      next: (data) => (this.userProfile = data)
    });

    // 2. Following
    this.followService.getFollowing(userId).subscribe({
      next: (ids) => {
        ids.forEach((id) =>
          this.profileService.getProfile(id).subscribe({
            next: (u) => this.following.push(u)
          })
        );
      }
    });

    // 3. Followers
    this.followService.getFollowers(userId).subscribe({
      next: (ids) => {
        ids.forEach((id) =>
          this.profileService.getProfile(id).subscribe({
            next: (u) => this.followers.push(u)
          })
        );
      }
    });

    // 4. Recommendations
    this.followService.getRecommendations(userId, 5).subscribe({
      next: (recommendations: RecommendationResponse[]) => {
        if (!recommendations || recommendations.length === 0) return;

        const requests = recommendations.map((rec) =>
          this.profileService.getProfile(rec.user_id)
        );
        forkJoin(requests).subscribe({
          next: (users) => (this.recommendations = users),
          error: (err) => console.error('Error fetching recommendations', err),
        });
      },
      error: (err) => console.error('Error fetching recommendations', err),
    });
  }

  onImgError(event: Event) {
    const img = event.target as HTMLImageElement;
    // ukloni handler da izbegnemo beskonačan loop ako i fallback padne
    img.onerror = null;
    img.src = this.DEFAULT_AVATAR;
  }


 }