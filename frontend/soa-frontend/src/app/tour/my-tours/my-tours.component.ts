import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TourService } from '../service/tour-service.service';


@Component({
  selector: 'app-my-tours',
  templateUrl: './my-tours.component.html',
  styleUrls: ['./my-tours.component.css']
})
export class MyToursComponent implements OnInit {

  tours: any[] = [];

  constructor(private tourService: TourService, private router: Router) {}

  ngOnInit(): void {
    this.loadMyTours();
  }

 loadMyTours() {
  const userStr = localStorage.getItem('user');
  if (!userStr) return;

  const user = JSON.parse(userStr);
  const authorId = user.id; // ili kako god se zove ID u user objektu

  this.tourService.getToursByAuthor(authorId).subscribe({
    next: (res) => this.tours = res,
    error: (err) => console.error(err)
  });
}

  createTour() {
    this.router.navigate(['home/create-tour']); // vodi na formu za kreiranje ture
  }

  addKeyPoints(tourId: string) {
    this.router.navigate(['tour-keypoints', tourId]);
  }

}
