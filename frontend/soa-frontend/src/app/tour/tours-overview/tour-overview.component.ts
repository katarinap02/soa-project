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

      this.tours = data.filter(t => t.status?.toUpperCase() === 'PUBLISHED');

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
