import { Component, Input, OnInit } from '@angular/core';
import { BlogPost } from '../model/BlogPost.model';
import { BlogService } from '../blog.service';
import { Route, Router } from '@angular/router';
import { UserService } from 'src/app/stakeholders/user.service';
import { FollowService } from 'src/app/followers/follow.service';
import { UserView } from 'src/app/stakeholders/model/UserView.model';
import { User } from 'src/app/stakeholders/model/User.model';

@Component({
  selector: 'app-view-blog-posts',
  templateUrl: './view-blog-posts.component.html',
  styleUrls: ['./view-blog-posts.component.css']
})
export class ViewBlogPostsComponent implements OnInit {

    @Input() posts: BlogPost[] = [];
    loggedUser!: UserView;
    constructor(private blogService: BlogService, private router: Router, private userService: UserService, private followService: FollowService) {}

    ngOnInit(): void {
    const userStr = localStorage.getItem('user');
    if (userStr) {
      const user = JSON.parse(userStr);
      this.loggedUser = user;
     console.log(this.loggedUser.username);
          
    }

    this.blogService.getAllBlogPosts().subscribe({
      next: (data) => {
        this.posts = [];

        data.forEach((post) => {
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

}
