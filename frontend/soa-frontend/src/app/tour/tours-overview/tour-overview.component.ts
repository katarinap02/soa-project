import { Component, OnInit } from '@angular/core';
import { Tour } from '../model/tour.model';
import { TourService } from '../service/tour-service.service';
import { ReviewComponent } from '../review/review.component';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ShoppingCartService } from '../service/shopping-cart.service';
import { ShoppingRpcService } from '../service/rpc-shopping-cart.service';


@Component({
  selector: 'app-tour-overview',
  templateUrl: './tour-overview.component.html',
  styleUrls: ['./tour-overview.component.css']
})
export class ToursOverviewComponent implements OnInit {
  tours: Tour[] = [];

  constructor(private tourService: TourService, 
    private dialog: MatDialog, 
    private router: Router,
    private shoppingCartService: ShoppingCartService,
    private shoppingRpcService: ShoppingRpcService) { }

  ngOnInit(): void {
    this.tourService.getAllTours().subscribe({
      next: data => this.tours = data,
      error: err => console.error(err)
    });
 
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
