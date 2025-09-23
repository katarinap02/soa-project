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
    status: 'draft',
    durations: {}
  };

  difficultyOptions: string[] = ['Lak', 'Srednji', 'Težak'];
  availableTags: string[] = ['Planinarenje', 'Grad', 'Istorija', 'Avantura', 'Priroda', 'Kultura'];

  selectedDifficulty: string = '';
  selectedTags: string[] = [];


  transportTypes: string[] = ['Walking', 'Bicycle', 'Car'];
  durations: { [key: string]: number | null } = {
    Walking: null,
    Bicycle: null,
    Car: null
  };


  constructor(private tourService: TourService) { }

createTour() {
  const userStr = localStorage.getItem('user');
  if (!userStr) return;

  const user = JSON.parse(userStr);
  const authorId = user.id;

  this.tour.tags = this.selectedTags;
  this.tour.difficulty = this.selectedDifficulty;
  this.tour.authorId = authorId;

  // --- LOGOVANJE ---
  console.log('User object from localStorage:', user);
  console.log('Author ID:', authorId);
  console.log('Tour object before sending:', this.tour);
  console.log('Selected tags:', this.selectedTags);
  console.log('Selected difficulty:', this.selectedDifficulty);
  // ------------------

  this.tour.durations = {
    walking: this.durations['Walking'] ?? undefined,
    bicycle: this.durations['Bicycle'] ?? undefined,
    car: this.durations['Car'] ?? undefined
  };

  this.tourService.createTour(this.tour).subscribe({
    next: res => {
      console.log('Response from backend:', res);
      alert('Tour created successfully');
      this.tour = { name: '', description: '', price: 0, difficulty: '', tags: [], status: 'draft' };
      this.selectedTags = [];
      this.selectedDifficulty = '';
    },
    error: err => {
      console.error('Error creating tour:', err);
      alert('Error creating tour');
    }
  });

 /*  this.tourService.createTourSaga(this.tour).subscribe({
    next: res => {
      console.log('Response from backend:', res);
      alert('Tour created successfully');
      this.tour = { name: '', description: '', price: 0, difficulty: '', tags: [], status: 'draft' };
      this.selectedTags = [];
      this.selectedDifficulty = '';
    },
    error: err => {
      console.error('Error creating tour:', err);
      alert('Error creating tour');
    }
  });*/
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
