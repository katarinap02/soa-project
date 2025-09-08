import { Component, OnInit } from '@angular/core';
import { Tour } from '../model/tour.model';
import { TourService } from '../service/tour-service.service';


@Component({
  selector: 'app-tour-overview',
  templateUrl: './tour-overview.component.html',
  styleUrls: ['./tour-overview.component.css']
})
export class ToursOverviewComponent implements OnInit {
  tours: Tour[] = [];

  constructor(private tourService: TourService) { }

  ngOnInit(): void {
    this.tourService.getAllTours().subscribe({
      next: data => this.tours = data,
      error: err => console.error(err)
    });
  }
}
