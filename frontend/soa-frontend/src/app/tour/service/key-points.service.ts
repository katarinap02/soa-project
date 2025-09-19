import { Injectable } from '@angular/core';
import {  HttpClient, HttpParams  } from '@angular/common/http';
import { Observable } from 'rxjs';
import { KeyPoint } from '../model/keyPoint.model';



@Injectable({
  providedIn: 'root'
})
export class KeyPointService {
  private apiUrl = 'http://localhost:8082/keypoints';

  constructor(private http: HttpClient) { }

  addKeyPoint(keyPoint: KeyPoint): Observable<any> {
    return this.http.post(this.apiUrl, keyPoint);
  }

    getKeyPointsByTour(tourId: string): Observable<KeyPoint[]> {
    const params = new HttpParams().set('tourId', tourId);
    return this.http.get<KeyPoint[]>(`${this.apiUrl}/by-tour`, { params });
  }

  
  updateKeyPoint(id: string, kp: KeyPoint): Observable<any> {
    return this.http.put(`${this.apiUrl}?id=${id}`, kp);
  }

  
  deleteKeyPoint(id: string): Observable<any> {
    return this.http.delete(`${this.apiUrl}?id=${id}`);
  }
}
