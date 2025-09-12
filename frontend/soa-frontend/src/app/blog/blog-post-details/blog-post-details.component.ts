import { Component, Input, OnInit } from '@angular/core';
import { BlogPost } from '../model/BlogPost.model';
import { ActivatedRoute, Router } from '@angular/router';
import { BlogService } from '../blog.service';
import { Comment } from '../model/Comment.model';

@Component({
  selector: 'app-blog-post-details',
  templateUrl: './blog-post-details.component.html',
  styleUrls: ['./blog-post-details.component.css']
})
export class BlogPostDetailsComponent implements OnInit {

   post!: BlogPost;
     @Input() comments: Comment[] = [];

  constructor(private route: ActivatedRoute, private blogService: BlogService, private router: Router) {}

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
  this.blogService.getBlogPostById(id).subscribe({
  next: (data) => {
    this.post = data;

    this.blogService.getCommentsByPostId(id).subscribe({
      next: (commentsData) => {
        this.comments = commentsData;
      },
      error: (err) => console.error('Failed to load comments:', err)
    });
  },
  error: (err) => console.error('Failed to load post:', err)
});
    }
  }

    createComment(id?: string): void {
    this.router.navigate(['/create-comment', id]);
  }

}
