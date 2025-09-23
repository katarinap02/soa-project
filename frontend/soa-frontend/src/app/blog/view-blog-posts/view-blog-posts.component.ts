import { Component, Input, OnInit } from '@angular/core';
import { BlogPost } from '../model/BlogPost.model';
import { BlogService } from '../blog.service';
import { Route, Router } from '@angular/router';
import { UserService } from 'src/app/stakeholders/user.service';
import { FollowService } from 'src/app/followers/follow.service';
import { UserView } from 'src/app/stakeholders/model/UserView.model';

@Component({
  selector: 'app-view-blog-posts',
  templateUrl: './view-blog-posts.component.html',
  styleUrls: ['./view-blog-posts.component.css']
})
export class ViewBlogPostsComponent implements OnInit {

    @Input() posts: BlogPost[] = [];
    loggedUser!: UserView;
    username!: string;
    constructor(private blogService: BlogService, private router: Router, private userService: UserService, private followService: FollowService) {}

    ngOnInit(): void {
    const userStr = localStorage.getItem('user');
    if (userStr) {
      const user = JSON.parse(userStr);
      this.loggedUser = user;
      this.username = user.name;
    }

    this.blogService.getAllBlogPosts().subscribe({
      next: (data) => {
        this.posts = [];
        console.log(this.username);

        this.posts = data;
        console.log(this.posts)
        data.forEach((post) => {
          post.likesCount = post.Likes?.length || 0;
          post.likedByCurrentUser = post.Likes?.some((like: any) => like.username === this.username) || false;

          //moj post
          if (post.username === this.loggedUser.username) {
            this.posts.push(post);
          } else {
            // da li pratim autora
            this.userService.getUserByUsername(post.username).subscribe({
              next: (author) => {
                this.followService.isFollowing(this.loggedUser.id, author.id).subscribe({
                  next: (isFollowing) => {
                    console.log("Check:", this.loggedUser.id, "->", author.id, "=", isFollowing);
                    if (isFollowing) {
                      this.posts.push(post);
                    }
                  }
                });

              },
              error: (err) =>
                console.error(`Greška pri dohvatanju autora (${post.username})`, err),
            });
          }
        });
      },
      error: (err) => {
        console.error('Failed to load blog posts:', err);
      },
    });
  }

     goToDetails(id?: string): void {
    this.router.navigate(['/post-details', id]);
  }

  likeBlog(post: any): void {
    this.blogService.likeBlog(this.username, post.id).subscribe({
      next: () => {
        if (!post.likesCount) post.likesCount = 0;
        post.likesCount++;
        post.likedByCurrentUser = true;
      },
      error: err => console.error("Error liking post:", err)
    });
  }

  unlikeBlog(post: any): void {
    this.blogService.unlikeBlog(this.username, post.id).subscribe({
      next: () => {
        if (!post.likesCount) post.likesCount = 0;
        post.likesCount--;
        post.likedByCurrentUser = false;
      },
      error: err => console.error("Error unliking post:", err)
    });
  }

}
