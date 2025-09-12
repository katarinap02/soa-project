import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { RecommendationResponse } from './model/recommendation-response.model';

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

}

