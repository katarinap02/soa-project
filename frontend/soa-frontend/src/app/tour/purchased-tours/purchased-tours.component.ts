import { Component, OnInit } from '@angular/core';
import { ShoppingCartService } from '../service/shopping-cart.service';
import { Router } from '@angular/router';
import { UserView } from 'src/app/stakeholders/model/UserView.model';
import { KeyPointService } from '../service/key-points.service';
import { KeyPoint } from '../model/keyPoint.model';
import { TourExecution } from '../model/tourExecution.model';
import { TourExecutionService } from '../service/tour-execution.service';
import { ReviewDialogComponent } from '../review-dialog/review-dialog.component';
import { ReviewComponent } from '../review/review.component';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-purchased-tours',
  templateUrl: './purchased-tours.component.html',
  styleUrls: ['./purchased-tours.component.css']
})
export class PurchasedToursComponent implements OnInit {

  tours: any[] = [];
  user: UserView | null = null;

  constructor(
    private shoppingCartService: ShoppingCartService,
    private router: Router,
    private keyPointService: KeyPointService,
    private tourExecutionService: TourExecutionService,
    private dialog: MatDialog, 
  ) {}

  ngOnInit(): void {
    const user = JSON.parse(localStorage.getItem('user')!);
    if (!user) return;

    this.user = user

    this.shoppingCartService.getPurchasedTours(user.id).subscribe({
      next: data => this.tours = data,
      error: err => console.error(err)
    });
  }

  viewMap(tourId: string) {
    this.router.navigate(['home/view-map-tourist', tourId]);
  }

  startTour(tourId: string): void {
  if (!this.user || !tourId) {
      alert('You must be logged in as a tourist to start a tour.');
      return;
    }

  this.keyPointService.getKeyPointsByTour(tourId).subscribe({
    next: (keyPoints: KeyPoint[]) => {
      if (!keyPoints || keyPoints.length === 0) {
        alert('This tour cannot be started because it has no key points.');
        return;
      }

    const touristId = this.user!.id;
      this.tourExecutionService.getActiveToursByTourist(touristId).subscribe({
        next: (executions: TourExecution[]) => {
          const existing = executions.find(te => te.tourId === tourId && te.status === 'active');

          if (existing) {
            console.log("Already active tour execution:", existing);
           this.router.navigate(['home/tour-execution', existing.id]);
          } else {
            this.tourExecutionService.startTour(tourId, touristId).subscribe({
              next: (newExecution: TourExecution) => {
                console.log("Started new tour execution:", newExecution);
                this.router.navigate(['home/tour-execution', newExecution.id]);
              },
              error: err => {
                console.error("Error starting tour:", err);
                alert("Could not start the tour.");
              }
            });
          }
        },
        error: err => {
          console.error("Error fetching active tours:", err);
        }
      });
    },
    error: err => {
      console.error("Error fetching key points:", err);
      alert('Could not check tour key points.');
    }
  });
  }

    openReviews(tourId: string) {
    this.dialog.open(ReviewComponent, {
      width: '600px',
      data: { tourId }
    });}
  
}
