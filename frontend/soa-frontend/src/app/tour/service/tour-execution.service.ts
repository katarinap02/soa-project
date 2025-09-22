import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { TourExecution } from '../model/tourExecution.model';
import { CompletedKeyPoint } from '../model/completedKeyPoint.model';

@Injectable({
  providedIn: 'root'
})
export class TourExecutionService {
  private apiUrl = 'http://localhost:8085/tours/tour-executions';

  constructor(private http: HttpClient) {}

  startTour(tourId: string, touristId: string): Observable<TourExecution> {
    return this.http.post<TourExecution>(`${this.apiUrl}/start`, { tourId, touristId });
  }

  getActiveToursByTourist(touristId: string): Observable<TourExecution[]> {
    return this.http.get<TourExecution[]>(`${this.apiUrl}/active`, { params: { touristId } });
  }

  getTourExecution(id: string): Observable<TourExecution> {
    return this.http.get<TourExecution>(`${this.apiUrl}/${id}`);
  }


  updateActivity(id: string): Observable<TourExecution> {
  return this.http.put<TourExecution>(`${this.apiUrl}/${id}/activity`, null);
}


  completeTour(id: string): Observable<TourExecution> {
    return this.http.put<TourExecution>(`${this.apiUrl}/${id}/complete`, {});
  }

  abandonTour(id: string): Observable<TourExecution> {
    return this.http.put<TourExecution>(`${this.apiUrl}/${id}/abandon`, {});
  }

  completeKeyPoint(id: string, keyPointId: string): Observable<CompletedKeyPoint> {
    return this.http.post<CompletedKeyPoint>(`${this.apiUrl}/${id}/keypoints`, { keyPointId });
  }

  getCompletedKeyPoints(id: string): Observable<CompletedKeyPoint[]> {
    return this.http.get<CompletedKeyPoint[]>(`${this.apiUrl}/${id}/keypoints`);
  }
}
