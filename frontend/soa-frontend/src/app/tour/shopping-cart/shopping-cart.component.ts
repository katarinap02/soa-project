import { Component, OnInit } from '@angular/core';
import { ShoppingCartService } from '../service/shopping-cart.service';
import { Router } from '@angular/router';
import { ShoppingRpcService } from '../service/rpc-shopping-cart.service';

@Component({
  selector: 'app-shopping-cart',
  templateUrl: './shopping-cart.component.html',
  styleUrls: ['./shopping-cart.component.css']
})
export class ShoppingCartComponent implements OnInit {

  cartTours: any[] = [];
  totalPrice: number = 0;

  constructor(private shoppingRpcService: ShoppingRpcService,
    private shoppingCartService: ShoppingCartService, 
    private router: Router) { }

  ngOnInit(): void {
    const user = JSON.parse(localStorage.getItem('user')!);
    if (!user) return;

    this.loadCart(user.id);
  }

loadCart(userId: string) {
  this.shoppingCartService.getCart(userId).subscribe({
     next: (data: any) => {
      console.log('Cart data:', data);
      
      this.cartTours = Array.isArray(data.items) ? data.items : [];
      
      this.totalPrice = this.cartTours.reduce((sum, t) => sum + t.price, 0);
    },
    error: err => console.error(err)
  });
}


  async checkout() {
    const user = JSON.parse(localStorage.getItem('user')!);
    if (!user) return;

    try {
      await this.shoppingRpcService.checkout(user.id);
      alert('Purchase successful!');
      this.cartTours = [];
      this.totalPrice = 0;
      this.router.navigate(['/home/purchased-tours']);
    } catch (err) {
      console.error(err);
      alert('Checkout failed!');
    }
  }

  removeFromCart(tourId: string) {
  const user = JSON.parse(localStorage.getItem('user')!);
  if (!user) return;

  // pozovi backend
  this.shoppingCartService.removeFromCart(user.id, tourId).subscribe({
    next: (updatedCart: any) => {
      this.cartTours = Array.isArray(updatedCart.items) ? updatedCart.items : [];
      this.totalPrice = this.cartTours.reduce((sum, t) => sum + t.price, 0);
    },
    error: err => console.error(err)
  });
}

}
