import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Review } from '../model/review.model';
import { ReviewService } from '../service/review.service';


@Component({
  selector: 'app-review-dialog',
  templateUrl: './review-dialog.component.html',
  styleUrls: ['./review-dialog.component.css']
})
export class ReviewDialogComponent {
  review: Review = {
    tourId: '',
    touristId: '',
    rating: 5,
    comment: '',
    visitDate: new Date(),
    commentDate: new Date(),
    images: []
  };

  constructor(
    public dialogRef: MatDialogRef<ReviewDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { tourId: string },
    private reviewService: ReviewService
  ) {
    this.review.tourId = data.tourId;
    const userStr = localStorage.getItem('user');
    if (userStr) {
      const user = JSON.parse(userStr);
      this.review.touristId = user.id;
    }
  }

  submit() {
    this.reviewService.addReview(this.review).subscribe({
      next: () => this.dialogRef.close(),
      error: err => console.error(err)
    });
  }
}
