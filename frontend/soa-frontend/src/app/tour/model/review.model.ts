export interface Review {
  id?: string;
  tourId: string;
  touristId: string;
  rating: number; // 1-5
  comment: string;
  visitDate: Date; 
  commentDate: Date; 
  images?: string[]; // URL slike
}
