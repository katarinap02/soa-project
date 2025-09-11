import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface KeyPoint {
  tourId: string;
  name: string;
  description: string;
  latitude: number;
  longitude: number;
  imageUrl: string;
}

@Injectable({
  providedIn: 'root'
})
export class KeyPointService {
  private apiUrl = 'http://localhost:8082/keypoints';

  constructor(private http: HttpClient) { }

  addKeyPoint(keyPoint: KeyPoint): Observable<any> {
    return this.http.post(this.apiUrl, keyPoint);
  }
}
