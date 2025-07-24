import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BlogService } from '../blog.service';
import { BlogPost } from '../model/BlogPost.model';

@Component({
  selector: 'app-create-blogpost',
  templateUrl: './create-blogpost.component.html',
  styleUrls: ['./create-blogpost.component.css']
})
export class CreateBlogpostComponent implements OnInit {

  blogPostForm!: FormGroup;
  message = '';

  constructor(private fb: FormBuilder, private blogService: BlogService)
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

    let blogPostData : BlogPost = {
   
      Username: "marko123",
      Title: this.blogPostForm.value.title,
      Description: this.blogPostForm.value.description,
      Date: new Date().toISOString()
    }

  
      this.blogService.createPost(blogPostData).subscribe({
        next: () => this.message = 'Uspešno kreiran post!',
        error: err => this.message = 'Greška: ' + (err.error?.message || 'Nepoznata greška')
      });
  }

}
