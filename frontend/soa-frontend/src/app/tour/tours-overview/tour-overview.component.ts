import { Component, OnInit } from '@angular/core';
import { Tour } from '../model/tour.model';
import { TourService } from '../service/tour-service.service';
import { ReviewComponent } from '../review/review.component';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { UserView } from 'src/app/stakeholders/model/UserView.model';
import { TourExecutionService } from '../service/tour-execution.service';
import { TourExecution } from '../model/tourExecution.model';


@Component({
  selector: 'app-tour-overview',
  templateUrl: './tour-overview.component.html',
  styleUrls: ['./tour-overview.component.css']
})
export class ToursOverviewComponent implements OnInit {
  tours: Tour[] = [];
  user: UserView | null = null;

  constructor(private tourService: TourService, private dialog: MatDialog, private router: Router, private tourExecutionService: TourExecutionService) { }

  ngOnInit(): void {
  this.loadLoggedUser();
  this.tourService.getAllTours().subscribe({
    next: data => {
      // filtriraj sve ture osim one koja se zove "Beogradska Tura"
      this.tours = data.filter(tour => tour.name !== "Beogradska Tura");
    },
    error: err => console.error(err)
  });
}

  loadLoggedUser() {
    const token = localStorage.getItem('token');
    console.log(token)
    
    const userStr = localStorage.getItem('user');
    if (!token || !userStr) {
      alert('No user logged in');
      return;
    }
    const user = JSON.parse(userStr);
    this.user = user;
  }

  openReviews(tourId: string) {
  this.dialog.open(ReviewComponent, {
    width: '600px',
    data: { tourId }
  });}

      viewMap(tourId: string) {
    this.router.navigate(['home/view-map-tourist', tourId]);
  }

  startTour(tourId: string): void {
  if (!this.user || !tourId) {
      alert('You must be logged in as a tourist to start a tour.');
      return;
    }

    const touristId = this.user.id;
    this.tourExecutionService.getActiveToursByTourist(touristId).subscribe({
      next: (executions: TourExecution[]) => {
        const existing = executions.find(te => te.tourId === tourId && te.status === 'active');

        if (existing) {
          console.log("Already active tour execution:", existing);
          //this.router.navigate(['home/view-map-tourist', existing.tourId]);
        } else {
          // ako ne postoji, kreiraj novu
          this.tourExecutionService.startTour(tourId, touristId).subscribe({
            next: (newExecution: TourExecution) => {
              console.log("Started new tour execution:", newExecution);
             // this.router.navigate(['home/view-map-tourist', newExecution.tourId]);
            },
            error: err => {
              console.error("Error starting tour:", err);
              alert("Could not start the tour.");
            }
          });
        }
      },
      error: err => {
        console.error("Error fetching active tours:", err);
      }
    });
  }


}
