import { Injectable } from '@angular/core';
import { AddToCartRequest, CheckoutRequest, CheckoutResponse, ShoppingCartResponse } from 'src/app/grpc/shopping_pb';
import { ShoppingCartServiceClient } from 'src/app/grpc/ShoppingServiceClientPb';

@Injectable({
  providedIn: 'root'
})
export class ShoppingRpcService {
  private client: ShoppingCartServiceClient;

  constructor() {
    // URL gde gRPC-Web proxy sluša
    this.client = new ShoppingCartServiceClient('http://localhost:8087', null, null);
  }

  addToCart(userId: string, tourId: string): Promise<ShoppingCartResponse> {
    return new Promise((resolve, reject) => {
      const req = new AddToCartRequest();
      req.setUserid(userId);
      req.setTourid(tourId);

      this.client.addToCart(req, {}, (err, response) => {
        if (err) reject(err);
        else resolve(response);
      });
    });
  }

  checkout(userId: string): Promise<CheckoutResponse> {
    return new Promise((resolve, reject) => {
      const req = new CheckoutRequest();
      req.setUserid(userId);

      this.client.checkout(req, {}, (err, response) => {
        if (err) reject(err);
        else resolve(response);
      });
    });
  }
}
