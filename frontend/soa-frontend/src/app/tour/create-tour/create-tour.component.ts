import { Component } from '@angular/core';
import { Tour } from '../model/tour.model';
import { TourService } from '../service/tour-service.service';


@Component({
  selector: 'app-create-tour',
  templateUrl: './create-tour.component.html',
  styleUrls: ['./create-tour.component.css']
})
export class CreateTourComponent {
  tour: Tour = {
    name: '',
    description: '',
    price: 0,
    weight: '',
    tags: [],
    status: 'draft'
  };
   tagsString: string = '';

  authorId = '64f8f3a2b5e4c8d1a2f1b9c0'; // primer, kasnije može iz login-a

  constructor(private tourService: TourService) { }

  createTour() {
    this.tourService.createTour(this.tour, this.authorId).subscribe({
      next: res => {
        alert('Tour created successfully');
        this.tour = { name: '', description: '', price: 0, weight: '', tags: [], status: 'draft' };
      },
      error: err => {
        console.error(err);
        alert('Error creating tour');
      }
    });
  }
}
