import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BlogService } from '../blog.service';
import { Comment } from '../model/Comment.model';

@Component({
  selector: 'app-create-comment',
  templateUrl: './create-comment.component.html',
  styleUrls: ['./create-comment.component.css']
})
export class CreateCommentComponent implements OnInit {
   commentForm!: FormGroup;
    message = '';
  
    constructor(private fb: FormBuilder, private blogService: BlogService)
    {}
    
    ngOnInit() {
      this.commentForm = this.fb.group({
        text: ['', Validators.required]
    
      });
    }
  
    onSubmit()
    {
      if(this.commentForm.invalid) return;
  
      let commentData : Comment = {  
        Username: "petar123",
        Text: this.commentForm.value.text,
        PostId: "4663d4c2-8f26-4ce0-8172-677abc51d0b8", //zakucano za sad
        DateCreated: new Date().toISOString(),
        DateModified: new Date().toISOString()
      }
  
    
        this.blogService.createComment(commentData).subscribe({
          next: () => this.message = 'Uspešno kreiran post!',
          error: err => this.message = 'Greška: ' + (err.error?.message || 'Nepoznata greška')
        });
    }

}
