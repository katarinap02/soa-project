import { Injectable } from '@angular/core';
import { BlogPost } from './model/BlogPost.model';
import { Observable } from 'rxjs';
import { HttpBackend, HttpClient } from '@angular/common/http';
import { Comment } from './model/Comment.model';

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

  createComment(comment: Comment): Observable<any>
  {
    return this.http.post<any>(this.apiUrl + '/create-comment', comment);
  }

   getAllBlogPosts(): Observable<BlogPost[]> {
    return this.http.get<BlogPost[]>(this.apiUrl);
  }

  getBlogPostsByUsername(username: string): Observable<BlogPost[]> {
    return this.http.get<BlogPost[]>(`${this.apiUrl}/by-username?username=${username}`);
  }

  getBlogPostById(id: string): Observable<BlogPost> {
      return this.http.get<BlogPost>(`${this.apiUrl}/by-id?id=${id}`);
}

getCommentsByPostId(postId: string): Observable<Comment[]> {
  return this.http.get<Comment[]>(`${this.apiUrl}/comments?postId=${postId}`);
}


}
