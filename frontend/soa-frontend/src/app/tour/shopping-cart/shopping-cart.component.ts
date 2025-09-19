import { Component, OnInit } from '@angular/core';
import { ShoppingCartService } from '../service/shopping-cart.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-shopping-cart',
  templateUrl: './shopping-cart.component.html',
  styleUrls: ['./shopping-cart.component.css']
})
export class ShoppingCartComponent implements OnInit {

  cartTours: any[] = [];
  totalPrice: number = 0;

  constructor(private shoppingCartService: ShoppingCartService, private router: Router) { }

  ngOnInit(): void {
    const user = JSON.parse(localStorage.getItem('user')!);
    if (!user) return;

    this.loadCart(user.id);
  }

  loadCart(userId: string) {
    this.shoppingCartService.getCart(userId).subscribe({
      next: data => {
        this.cartTours = data;
        this.totalPrice = this.cartTours.reduce((sum, t) => sum + t.price, 0);
      },
      error: err => console.error(err)
    });
  }

  checkout() {
    const user = JSON.parse(localStorage.getItem('user')!);
    if (!user) return;

    this.shoppingCartService.checkout(user.id).subscribe({
      next: () => {
        alert('Purchase successful!');
        this.cartTours = [];
        this.totalPrice = 0;
        this.router.navigate(['/home/purchased-tours']);
      },
      error: err => console.error(err)
    });
  }
}
