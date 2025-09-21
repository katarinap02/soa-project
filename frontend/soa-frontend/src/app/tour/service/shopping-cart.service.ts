import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ShoppingCartService {

  private apiUrl = 'http://localhost:8082/cart';

  constructor(private http: HttpClient) { }


  getPurchasedTours(userId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/purchased?userId=${userId}`);
  }

  getCart(userId: string) {
  return this.http.get<any[]>(`${this.apiUrl}?userId=${userId}`);
}

removeFromCart(userId: string, tourId: string) {
  return this.http.delete(`${this.apiUrl}/remove?userId=${userId}&tourId=${tourId}`);
}



}
