import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { UserView } from 'src/app/stakeholders/model/UserView.model';
import { Profile } from '../model/profile.model';
import { FollowService } from '../follow.service';
import { UserService } from 'src/app/stakeholders/user.service';
import { RecommendationResponse } from '../model/recommendation-response.model';
import { forkJoin, map } from 'rxjs';

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

  editMode = false;
 
  editProfile: Profile = {
  id: '',
  firstName: '',
  lastName: '',
  profilePicture: '',
  biography: '',
  motto: ''
};


  constructor(
    private profileService: UserService,
    private followService: FollowService
  ) {}

  ngOnInit(): void {
  const userStr = localStorage.getItem('user');
  if (!userStr) return;

  const user = JSON.parse(userStr);
  const userId = user.id;

  this.profileService.getProfile(userId).subscribe({
    next: (data) => {
      this.userProfile = { ...data, id: userId }; // dodaj id
       console.log('Id:', userId);
      console.log('Fetched profile:', data);

      this.refreshAllData();
    },
    error: (err) => console.error('Error fetching user profile', err)
  });
}

refreshAllData() {
  if (!this.userProfile) return;
  const userId = this.userProfile.id;

  // --- Following ---
  this.followService.getFollowing(userId).subscribe({
    next: (ids) => {
      const requests = ids.map(id =>
        this.profileService.getProfile(id).pipe(map(profile => ({ ...profile, id })))
      );
      if (requests.length > 0) {
        forkJoin(requests).subscribe(users => this.following = users);
      } else {
        this.following = [];
      }
    },
    error: (err) => console.error('Error fetching following', err)
  });

  // --- Followers ---
  this.followService.getFollowers(userId).subscribe({
    next: (ids) => {
      const requests = ids.map(id =>
        this.profileService.getProfile(id).pipe(map(profile => ({ ...profile, id })))
      );
      if (requests.length > 0) {
        forkJoin(requests).subscribe(users => this.followers = users);
      } else {
        this.followers = [];
      }
    },
    error: (err) => console.error('Error fetching followers', err)
  });

  // --- Recommendations ---
  this.followService.getRecommendations(userId, 5).subscribe({
    next: (recs) => {
      const validRecs = recs.filter(r => r.user_id && r.user_id !== '');
      const requests = validRecs.map(r =>
        this.profileService.getProfile(r.user_id).pipe(map(profile => ({ ...profile, id: r.user_id })))
      );
      if (requests.length > 0) {
        forkJoin(requests).subscribe(users => this.recommendations = users);
      } else {
        this.recommendations = [];
      }
    },
    error: (err) => console.error('Error fetching recommendations', err)
  });
}

  onImgError(event: Event) {
    const img = event.target as HTMLImageElement;
    img.onerror = null;
    img.src = this.DEFAULT_AVATAR;
  }

  toggleFollow(target: Profile) {
  if (!this.userProfile) return;

  const followerId = this.userProfile.id;
  const followeeId = target.id;

  if (this.isFollowing(followeeId)) {
    // Unfollow
    this.followService.unfollowUser(followerId, followeeId).subscribe({
      next: () => this.refreshAllData(),
      error: (err) => console.error('Unfollow failed', err)
    });
  } else {
    // Follow
    this.followService.followUser(followerId, followeeId).subscribe({
      next: () => this.refreshAllData(),
      error: (err) => console.error('Follow failed', err)
    });
  }
}

// Helper funkcija
isFollowing(userId: string): boolean {
  return this.following.some(f => f.id === userId);
}


toggleEdit() {
  this.editMode = !this.editMode;
  if (this.editMode && this.userProfile) {
    // pravi kopiju da ne menjaš original odmah
    this.editProfile = { ...this.userProfile };
  }
}

cancelEdit() {
  this.editMode = false;
  this.editProfile = {
    id: '',
    firstName: '',
    lastName: '',
    profilePicture: '',
    biography: '',
    motto: ''
  };
}


saveChanges() {
  if (!this.editProfile) return;
  const userId = this.userProfile?.id;
  if (!userId) return;

  this.profileService.updateProfile(userId, this.editProfile).subscribe({
    next: () => {
      this.userProfile = { ...this.editProfile! };
      this.editMode = false;

      // ✅ Običan alert
      alert('Profile updated successfully');
    },
    error: (err) => {
      console.error('Update failed', err);
      alert('Failed to update profile');
    },
  });
}

 }