import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from './model/User.model';
import { UserView } from './model/UserView.model';
import { Profile } from '../followers/model/profile.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = 'http://localhost:8080/users'; 
  private apiUrl1 = 'http://localhost:8080';
  

  constructor(private http: HttpClient) {}

  register(user: User): Observable<User> {
    return this.http.post<User>(`${this.apiUrl}/register`, user);
  }

  getAllUsers(): Observable<UserView[]> {
    return this.http.get<UserView[]>(this.apiUrl);
  }

  getProfile(userId: string): Observable<Profile> {
    return this.http.get<Profile>(`${this.apiUrl1}/profile/${userId}`);
  }

  getUserByUsername(username: string): Observable<UserView> {
  return this.http.post<UserView>(`${this.apiUrl}/by-username`, { username });
}

}
