import { Component, OnInit } from '@angular/core';
import { ShoppingCartService } from '../service/shopping-cart.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-purchased-tours',
  templateUrl: './purchased-tours.component.html',
  styleUrls: ['./purchased-tours.component.css']
})
export class PurchasedToursComponent implements OnInit {

  tours: any[] = [];

  constructor(
    private shoppingCartService: ShoppingCartService,
    private router: Router
  ) {}

  ngOnInit(): void {
    const user = JSON.parse(localStorage.getItem('user')!);
    if (!user) return;

    this.shoppingCartService.getPurchasedTours(user.id).subscribe({
      next: data => this.tours = data,
      error: err => console.error(err)
    });
  }

  viewMap(tourId: string) {
    this.router.navigate(['home/view-map-tourist', tourId]);
  }
}
