import { Component, OnInit } from '@angular/core';
import { Tour } from '../model/tour.model';
import { TourService } from '../service/tour-service.service';
import { ReviewComponent } from '../review/review.component';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { UserView } from 'src/app/stakeholders/model/UserView.model';
import { TourExecutionService } from '../service/tour-execution.service';
import { TourExecution } from '../model/tourExecution.model';
import { KeyPoint } from '../model/keyPoint.model';
import { KeyPointService } from '../service/key-points.service';
import { ShoppingCartService } from '../service/shopping-cart.service';
import { ShoppingRpcService } from '../service/rpc-shopping-cart.service';


@Component({
  selector: 'app-tour-overview',
  templateUrl: './tour-overview.component.html',
  styleUrls: ['./tour-overview.component.css']
})
export class ToursOverviewComponent implements OnInit {
  tours: Tour[] = [];
  user: UserView | null = null;

  constructor(private tourService: TourService,
     private dialog: MatDialog, 
     private router: Router, 
     private tourExecutionService: TourExecutionService, 
     private keyPointService: KeyPointService,
     private shoppingCartService: ShoppingCartService,
     private shoppingRpcService: ShoppingRpcService) { }


ngOnInit(): void {
   this.loadLoggedUser();
  this.tourService.getAllTours().subscribe({
    next: data => {
    
      this.tours = data.filter(t => t.status?.toUpperCase() !== 'ARCHIVED');
      this.tours = data.filter(tour => tour.name !== "Beogradska Tura");

    },
    error: err => console.error(err)
  });
}

  loadLoggedUser() {
    const token = localStorage.getItem('token');
    console.log(token)
    
    const userStr = localStorage.getItem('user');
    if (!token || !userStr) {
      alert('No user logged in');
      return;
    }
    const user = JSON.parse(userStr);
    this.user = user;
  }


  openReviews(tourId: string) {
  this.dialog.open(ReviewComponent, {
    width: '600px',
    data: { tourId }
  });}

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


async buyTour(tourId?: string) {
    if (!tourId) {
      alert('Tour ID is missing!');
      return;
    }

    const user = JSON.parse(localStorage.getItem('user')!);
    if (!user) {
      alert('Please login first');
      return;
    }


    try {
      await this.shoppingRpcService.addToCart(user.id, tourId);
      alert('Tour added to cart!');
    } catch (err) {
      console.error(err);
      alert('Failed to add tour to cart.');
    }
  }


goToCart(): void {
  this.router.navigate(['/home/shopping-cart']);
}

}
