import { Component, Input, OnInit } from '@angular/core';
import { BlogPost } from '../model/BlogPost.model';
import { BlogService } from '../blog.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-my-blog-posts',
  templateUrl: './my-blog-posts.component.html',
  styleUrls: ['./my-blog-posts.component.css']
})
export class MyBlogPostsComponent implements OnInit {

    @Input() posts: BlogPost[] = [];
  
      constructor(private blogService: BlogService, private router: Router) {}
  
      ngOnInit(): void {

          const currentUser = localStorage.getItem('user');
       if (currentUser) {
            const userObj = JSON.parse(currentUser);   
          const username = userObj.username;    
          
          this.blogService.getBlogPostsByUsername(username).subscribe({
        next: (data) => {
     
          this.posts = data;
         
       
        
        },
        error: (err) => {
          console.error('Failed to load blog posts:', err);
        }
      });
        
      }
                                   
      }
   goToDetails(id?: string): void {
    this.router.navigate(['/post-details', id]);
  }
         

}
