import { Component, Input, OnInit } from '@angular/core';
import { BlogPost } from '../model/BlogPost.model';
import { BlogService } from '../blog.service';

@Component({
  selector: 'app-view-blog-posts',
  templateUrl: './view-blog-posts.component.html',
  styleUrls: ['./view-blog-posts.component.css']
})
export class ViewBlogPostsComponent implements OnInit {

    @Input() posts: BlogPost[] = [];

    constructor(private blogService: BlogService) {}

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

}
