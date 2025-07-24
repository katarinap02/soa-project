import { Injectable } from '@angular/core';
import { BlogPost } from './model/BlogPost.model';
import { Observable } from 'rxjs';
import { HttpBackend, HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class BlogService {

  private apiUrl = '/blog'

  constructor(private http: HttpClient) { }

  createPost(post: BlogPost): Observable<any>
  {
    return this.http.post<any>(this.apiUrl + '/create-post', post);
  }
}
