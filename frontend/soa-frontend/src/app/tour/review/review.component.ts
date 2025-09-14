import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog, MatDialogRef } from '@angular/material/dialog';
import { Review } from '../model/review.model';
import { ReviewService } from '../service/review.service';
import { ReviewDialogComponent } from '../review-dialog/review-dialog.component';


@Component({
  selector: 'app-review',
  templateUrl: './review.component.html',
  styleUrls: ['./review.component.css']
})
export class ReviewComponent implements OnInit {
  reviews: Review[] = [];

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: { tourId: string },
    private reviewService: ReviewService,
    private dialogRef: MatDialogRef<ReviewComponent>,
    private dialog: MatDialog
  ) {}

  ngOnInit(): void {
    this.loadReviews();
  }

  loadReviews() {
    this.reviewService.getReviewsByTour(this.data.tourId).subscribe({
      next: (res) => this.reviews = res,
      error: (err) => console.error(err)
    });
  }

  addReview() {
    const dialogRef = this.dialog.open(ReviewDialogComponent, {
      width: '400px',
      data: { tourId: this.data.tourId }
    });

    dialogRef.afterClosed().subscribe(() => {
      this.loadReviews(); // osveži listu recenzija
    });
  }
}