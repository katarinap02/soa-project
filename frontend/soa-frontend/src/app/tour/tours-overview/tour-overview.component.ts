import { Component, OnInit } from '@angular/core';
import { Tour } from '../model/tour.model';
import { TourService } from '../service/tour-service.service';
import { ReviewComponent } from '../review/review.component';
import { MatDialog } from '@angular/material/dialog';


@Component({
  selector: 'app-tour-overview',
  templateUrl: './tour-overview.component.html',
  styleUrls: ['./tour-overview.component.css']
})
export class ToursOverviewComponent implements OnInit {
  tours: Tour[] = [];

  constructor(private tourService: TourService, private dialog: MatDialog) { }

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
}
