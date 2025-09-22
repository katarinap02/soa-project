import { Component, Inject, OnInit, ViewEncapsulation } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog, MatDialogRef } from '@angular/material/dialog';
import { Review } from '../model/review.model';
import { ReviewService } from '../service/review.service';
import { ReviewComponent } from '../review/review.component';



@Component({
  selector: 'app-review-dialog',
  templateUrl: './review-dialog.component.html',
  styleUrls: ['./review-dialog.component.css'],
   encapsulation: ViewEncapsulation.None
})
export class ReviewDialogComponent implements OnInit {
  reviews: Review[] = [];       // lista svih recenzija
  review: Review = {            // objekat za formu
    tourId: '',                 // popuni u ngOnInit
    touristId: '',              // možeš dohvatiti iz auth servisa
    rating: 0,
    comment: '',
    visitDate: new Date(),
    commentDate: new Date(),
    images: []
  };

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: { tourId: string },
    private reviewService: ReviewService,
    private dialogRef: MatDialogRef<ReviewComponent>,
    private dialog: MatDialog
  ) {}

  ngOnInit(): void {
    this.review.tourId = this.data.tourId; // setuj tourId na početku
    this.loadReviews();
  }

  loadReviews() {
    this.reviewService.getReviewsByTour(this.data.tourId).subscribe({
      next: (res) => this.reviews = res,
      error: (err) => console.error(err)
    });
  }

  submit() {
    this.review.commentDate = new Date(); // automatski setuj vreme komentara

    this.reviewService.addReview(this.review).subscribe({
      next: () => {
        this.loadReviews();
        this.resetForm();
        alert('Review successfully added!');
      },
       error: (err) => {
      console.error(err);
      alert('Error adding review.');
    }
      
    });
  }

  resetForm() {
    this.review = {
      tourId: this.data.tourId,
      touristId: '',
      rating: 0,
      comment: '',
      visitDate: new Date(),
      commentDate: new Date(),
      images: []
    };
  }

  addReview() {
    const dialogRef = this.dialog.open(ReviewDialogComponent, {
      width: '400px',
      data: { tourId: this.data.tourId }
    });

    dialogRef.afterClosed().subscribe(() => {
      this.loadReviews();
    });
  }

  onFileSelected(event: any) {
  const files: FileList = event.target.files;
  this.review.images = []; // resetuj prethodne slike
  for (let i = 0; i < files.length; i++) {
    const reader = new FileReader();
    reader.onload = (e: any) => {
      this.review.images!.push(e.target.result); // base64 string
    };
    reader.readAsDataURL(files[i]);
  }
}

}
