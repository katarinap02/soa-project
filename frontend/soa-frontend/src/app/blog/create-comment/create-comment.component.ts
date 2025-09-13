import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BlogService } from '../blog.service';
import { Comment } from '../model/Comment.model';
import { ActivatedRoute, Router } from '@angular/router';
import { FollowService } from 'src/app/followers/follow.service';
import { UserService } from 'src/app/stakeholders/user.service';

@Component({
  selector: 'app-create-comment',
  templateUrl: './create-comment.component.html',
  styleUrls: ['./create-comment.component.css']
})
export class CreateCommentComponent implements OnInit {
   commentForm!: FormGroup;
    message = '';
    canComment = false;
  
    constructor(
    private fb: FormBuilder, 
    private blogService: BlogService, 
    private route: ActivatedRoute, 
    private router: Router,
    private userService: UserService,
    private followService: FollowService
  ) {}
    
    ngOnInit() {
     
      this.commentForm = this.fb.group({
        text: ['', Validators.required]
    
      });
      this.checkCommentPermission();
    }

    checkCommentPermission() {
    const currentUser = localStorage.getItem('user');
    if (!currentUser) return;

    const userObj = JSON.parse(currentUser);
    const postId = this.route.snapshot.paramMap.get('id');
    if (!postId) return;

    // Dobij sve postove i nađi ovaj post
    this.blogService.getAllBlogPosts().subscribe({
      next: (posts) => {
        const post = posts.find(p => p.id === postId);
        if (!post) return;

        // Ako je moj post
        if (post.username === userObj.username) {
          this.canComment = true;
          return;
        }

        // Ako nije moj, proveri da li pratim autora
        this.userService.getUserByUsername(post.username).subscribe({
          next: (author) => {
            this.followService.isFollowing(userObj.id, author.id).subscribe({
              next: (response) => {
                this.canComment = response === true;
                if (!this.canComment) {
                  this.message = 'Možete komentarisati samo svoje postove ili postove osoba koje pratite.';
                }
              }
            });
          }
        });
      }
    });
  }
  
    onSubmit()
    {
      if(this.commentForm.invalid || !this.canComment) return;

        const currentUser = localStorage.getItem('user');
       if (currentUser) {
            const userObj = JSON.parse(currentUser);   
          const username = userObj.username;   
    const id = this.route.snapshot.paramMap.get('id');
      let commentData : Comment = {  
        
        username: username,
        text: this.commentForm.value.text,
        postId: id, //zakucano za sad
        date_created: new Date().toISOString(),
        date_modified: new Date().toISOString()
      }
  
      console.log(commentData)
    
        this.blogService.createComment(commentData).subscribe({
          next: () => this.router.navigate(['/post-details', id]),
          error: err => this.message = 'Greška: ' + (err.error?.message || 'Nepoznata greška')
        });
    }
  }
}
