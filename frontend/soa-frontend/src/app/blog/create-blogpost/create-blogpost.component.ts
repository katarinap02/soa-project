import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BlogService } from '../blog.service';
import { BlogPost } from '../model/BlogPost.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-blogpost',
  templateUrl: './create-blogpost.component.html',
  styleUrls: ['./create-blogpost.component.css']
})
export class CreateBlogpostComponent implements OnInit {

  blogPostForm!: FormGroup;
  message = '';

  constructor(private fb: FormBuilder, private blogService: BlogService, private router: Router)
  {}
  
  ngOnInit() {
    this.blogPostForm = this.fb.group({
      title: ['', Validators.required],
      description: ['', Validators.required]
    });
  }



  onSubmit()
  {
    if(this.blogPostForm.invalid) return;

       const currentUser = localStorage.getItem('user');
       if (currentUser) {
            const userObj = JSON.parse(currentUser);   
          const username = userObj.username;    
          

    let blogPostData : BlogPost = {
    
      username: username,
      title: this.blogPostForm.value.title,
      description: this.blogPostForm.value.description,
      date: new Date().toISOString()
    }
    

     console.log(blogPostData);
  
      this.blogService.createPost(blogPostData).subscribe({
        next: () =>  this.router.navigate(['/view-my-blogposts']),
        error: err => this.message = 'Greška: ' + (err.error?.message || 'Nepoznata greška')
      });
  }
}

}
