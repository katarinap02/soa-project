import { Component, Input, OnInit } from '@angular/core';
import { BlogService } from '../blog.service';
import { BlogPost } from '../model/BlogPost.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-view-posts-new',
  templateUrl: './view-posts-new.component.html',
  styleUrls: ['./view-posts-new.component.css']
})
export class ViewPostsNewComponent implements OnInit{

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
