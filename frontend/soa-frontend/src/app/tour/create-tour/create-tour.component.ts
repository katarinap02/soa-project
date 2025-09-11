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
    difficulty: '',
    tags: [],
    status: 'draft'
  };
  difficultyOptions: string[] = ['Lak', 'Srednji', 'Težak'];
  availableTags: string[] = ['Planinarenje', 'Grad', 'Istorija', 'Avantura', 'Priroda', 'Kultura'];

  selectedDifficulty: string = '';
  selectedTags: string[] = [];




  constructor(private tourService: TourService) { }
  
createTour() {
  const userStr = localStorage.getItem('user');
  if (!userStr) return;

  const user = JSON.parse(userStr);
  const authorId = user.id;

  this.tour.tags = this.selectedTags;
  this.tour.difficulty = this.selectedDifficulty; // ranije weight

  this.tourService.createTour(this.tour, authorId).subscribe({
    next: res => {
      alert('Tour created successfully');
      this.tour = { name: '', description: '', price: 0, difficulty: '', tags: [], status: 'draft' };
      this.selectedTags = [];
      this.selectedDifficulty = '';
    },
    error: err => {
      console.error(err);
      alert('Error creating tour');
    }
  });
}

onTagChange(event: any, tag: string) {
  if (event.checked) {
    if (!this.selectedTags.includes(tag)) {
      this.selectedTags.push(tag);
    }
  } else {
    this.selectedTags = this.selectedTags.filter(t => t !== tag);
  }
}

}
