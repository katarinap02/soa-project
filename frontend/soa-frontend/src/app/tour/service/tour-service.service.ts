import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Tour } from '../model/tour.model';


@Injectable({
  providedIn: 'root'
})
export class TourService {
  private apiUrl = 'http://localhost:8082/tours';
  //private apiUrl = 'http://localhost:8085/tours/tours';

  constructor(private http: HttpClient) { }

  getAllTours(): Observable<Tour[]> {
    return this.http.get<Tour[]>(this.apiUrl);
  }

  createTour(tour: Tour): Observable<any> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });
    return this.http.post(this.apiUrl, tour, { headers });
  }

  // getToursByAuthor(authorId: string): Observable<Tour[]> {
  //   return this.http.get<Tour[]>(`${this.apiUrl}/by-author?authorId=${authorId}`);
  // }

getToursByAuthor(authorId: string): Observable<Tour[]> {
  const params = new HttpParams().set('authorId', authorId);
  return this.http.get<Tour[]>(`${this.apiUrl}/by-author`, { params });
}

    getTourById(tourId: string): Observable<Tour> {
    return this.http.get<Tour>(`${this.apiUrl}/${tourId}`);
  }

}
