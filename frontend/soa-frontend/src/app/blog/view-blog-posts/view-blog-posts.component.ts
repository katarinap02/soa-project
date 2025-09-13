import { Component, Input, OnInit } from '@angular/core';
import { BlogPost } from '../model/BlogPost.model';
import { BlogService } from '../blog.service';
import { Route, Router } from '@angular/router';

@Component({
  selector: 'app-view-blog-posts',
  templateUrl: './view-blog-posts.component.html',
  styleUrls: ['./view-blog-posts.component.css']
})
export class ViewBlogPostsComponent implements OnInit {

    @Input() posts: BlogPost[] = [];

    constructor(private blogService: BlogService, private router: Router) {}

    ngOnInit(): void {

       this.blogService.getAllBlogPosts().subscribe({
      next: (data) => {
   
        this.posts = data;
       
     
      
      },
      error: (err) => {
        console.error('Failed to load blog posts:', err);
      }
    });
      
    }

     goToDetails(id?: string): void {
    this.router.navigate(['/post-details', id]);
  }

}
