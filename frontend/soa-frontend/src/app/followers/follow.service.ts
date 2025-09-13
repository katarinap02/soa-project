import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { RecommendationResponse } from './model/recommendation-response.model';
import { FollowRequest } from './model/follow-request.model';

@Injectable({
  providedIn: 'root'
})
export class FollowService {
  private apiUrl = 'http://localhost:8084';

  constructor(private http: HttpClient) {}

  getFollowing(userId: string): Observable<string[]> {
    return this.http.get<string[]>(`${this.apiUrl}/following/${userId}`);
  }

  getFollowers(userId: string): Observable<string[]> {
    return this.http.get<string[]>(`${this.apiUrl}/followers/${userId}`);
  }

  getRecommendations(userId: string, limit: number = 1): Observable<RecommendationResponse[]> {
  return this.http.get<RecommendationResponse[]>(`${this.apiUrl}/recommendations/${userId}/${limit}`);
}

followUser(followerId: string, followeeId: string): Observable<string> {
  const body: FollowRequest = { follower_id: followerId, followee_id: followeeId };
  return this.http.post<string>(
    `${this.apiUrl}/follow`,
    body,
    { headers: { 'Content-Type': 'application/json' } }
  );
}


  unfollowUser(followerId: string, followeeId: string): Observable<string> {
    const body: FollowRequest = { follower_id: followerId, followee_id: followeeId };
    return this.http.request<string>('delete', `${this.apiUrl}/unfollow`, { body });
  }

  isFollowing(followerId: string, followeeId: string): Observable<boolean> {
    return this.http.get<boolean>(`${this.apiUrl}/is-following/${followerId}/${followeeId}`);
  }

}

