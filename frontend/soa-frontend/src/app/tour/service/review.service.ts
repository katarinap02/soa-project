import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Review } from '../model/review';


@Injectable({
  providedIn: 'root'
})
export class ReviewService {
  private apiUrl = 'http://localhost:8082/reviews';

  constructor(private http: HttpClient) {}

  getReviewsByTour(tourId: string): Observable<Review[]> {
    return this.http.get<Review[]>(`${this.apiUrl}/by-tour?tourId=${tourId}`);
  }

  addReview(review: Review): Observable<any> {
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    return this.http.post(this.apiUrl, review, { headers });
  }
}
