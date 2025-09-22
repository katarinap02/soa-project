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
  filteredTours: any[] = [];

  showDraft = true;
  showPublished = true;
  showArchived = true;

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
    // next: (res) => this.tours = res,
    // error: (err) => console.error(err)
        next: (res) => {
      console.log("✅ Tours koje sam dobio sa backenda:", res);
      this.tours = res;
      this.applyFilter();
    },
    error: (err) => {
      console.error("❌ Greška sa backenda:", err);
    }
  });
}

  applyFilter() {
    this.filteredTours = this.tours.filter(tour =>
      (this.showDraft && tour.status === 'draft') ||
      (this.showPublished && tour.status === 'published') ||
      (this.showArchived && tour.status === 'archived')
    );
  }

  toggleFilter(status: string) {
    if (status === 'draft') this.showDraft = !this.showDraft;
    if (status === 'published') this.showPublished = !this.showPublished;
    if (status === 'archived') this.showArchived = !this.showArchived;

    this.applyFilter();
  }

  createTour() {
    this.router.navigate(['home/create-tour']); // vodi na formu za kreiranje ture
  }

  addKeyPoints(tourId: string) {
    this.router.navigate(['home/tour-keypoints', tourId]);
  }

    viewMap(tourId: string) {
    this.router.navigate(['home/view-map', tourId]);
  }


  changeStatus(tour: any, newStatus: string) {
    this.tourService.updateTourStatus(tour.id, newStatus).subscribe({
      next: (res) => {
        console.log(`Tour ${tour.id} updated to ${newStatus}`);
        tour.status = newStatus;
        this.applyFilter();
      },
      error: (err) => {
        console.error("Error updating status:", err);
      }
    });
  }


}
