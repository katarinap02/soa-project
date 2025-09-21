import * as jspb from 'google-protobuf'



export class AddToCartRequest extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): AddToCartRequest;

  getTourid(): string;
  setTourid(value: string): AddToCartRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddToCartRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddToCartRequest): AddToCartRequest.AsObject;
  static serializeBinaryToWriter(message: AddToCartRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddToCartRequest;
  static deserializeBinaryFromReader(message: AddToCartRequest, reader: jspb.BinaryReader): AddToCartRequest;
}

export namespace AddToCartRequest {
  export type AsObject = {
    userid: string,
    tourid: string,
  }
}

export class ShoppingCartResponse extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): ShoppingCartResponse;

  getItemsList(): Array<OrderItem>;
  setItemsList(value: Array<OrderItem>): ShoppingCartResponse;
  clearItemsList(): ShoppingCartResponse;
  addItems(value?: OrderItem, index?: number): OrderItem;

  getTotalprice(): number;
  setTotalprice(value: number): ShoppingCartResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ShoppingCartResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ShoppingCartResponse): ShoppingCartResponse.AsObject;
  static serializeBinaryToWriter(message: ShoppingCartResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ShoppingCartResponse;
  static deserializeBinaryFromReader(message: ShoppingCartResponse, reader: jspb.BinaryReader): ShoppingCartResponse;
}

export namespace ShoppingCartResponse {
  export type AsObject = {
    userid: string,
    itemsList: Array<OrderItem.AsObject>,
    totalprice: number,
  }
}

export class OrderItem extends jspb.Message {
  getTourid(): string;
  setTourid(value: string): OrderItem;

  getTourname(): string;
  setTourname(value: string): OrderItem;

  getPrice(): number;
  setPrice(value: number): OrderItem;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderItem.AsObject;
  static toObject(includeInstance: boolean, msg: OrderItem): OrderItem.AsObject;
  static serializeBinaryToWriter(message: OrderItem, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderItem;
  static deserializeBinaryFromReader(message: OrderItem, reader: jspb.BinaryReader): OrderItem;
}

export namespace OrderItem {
  export type AsObject = {
    tourid: string,
    tourname: string,
    price: number,
  }
}

export class CheckoutRequest extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): CheckoutRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CheckoutRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CheckoutRequest): CheckoutRequest.AsObject;
  static serializeBinaryToWriter(message: CheckoutRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CheckoutRequest;
  static deserializeBinaryFromReader(message: CheckoutRequest, reader: jspb.BinaryReader): CheckoutRequest;
}

export namespace CheckoutRequest {
  export type AsObject = {
    userid: string,
  }
}

export class CheckoutResponse extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): CheckoutResponse;

  getTokensList(): Array<TourPurchaseToken>;
  setTokensList(value: Array<TourPurchaseToken>): CheckoutResponse;
  clearTokensList(): CheckoutResponse;
  addTokens(value?: TourPurchaseToken, index?: number): TourPurchaseToken;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CheckoutResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CheckoutResponse): CheckoutResponse.AsObject;
  static serializeBinaryToWriter(message: CheckoutResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CheckoutResponse;
  static deserializeBinaryFromReader(message: CheckoutResponse, reader: jspb.BinaryReader): CheckoutResponse;
}

export namespace CheckoutResponse {
  export type AsObject = {
    userid: string,
    tokensList: Array<TourPurchaseToken.AsObject>,
  }
}

export class TourPurchaseToken extends jspb.Message {
  getTourid(): string;
  setTourid(value: string): TourPurchaseToken;

  getToken(): string;
  setToken(value: string): TourPurchaseToken;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TourPurchaseToken.AsObject;
  static toObject(includeInstance: boolean, msg: TourPurchaseToken): TourPurchaseToken.AsObject;
  static serializeBinaryToWriter(message: TourPurchaseToken, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TourPurchaseToken;
  static deserializeBinaryFromReader(message: TourPurchaseToken, reader: jspb.BinaryReader): TourPurchaseToken;
}

export namespace TourPurchaseToken {
  export type AsObject = {
    tourid: string,
    token: string,
  }
}

