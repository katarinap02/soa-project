import { AfterViewInit, ChangeDetectionStrategy, Component, OnDestroy, OnInit } from '@angular/core';
import * as L from 'leaflet';
import { TourExecution } from '../model/tourExecution.model';
import { ActivatedRoute } from '@angular/router';
import { TourExecutionService } from '../service/tour-execution.service';
import { KeyPoint } from '../model/keyPoint.model';
import { KeyPointService } from '../service/key-points.service';
import { TourService } from '../service/tour-service.service';
import { Tour } from '../model/tour.model';
import { CompletedKeyPoint } from '../model/completedKeyPoint.model';
import { forkJoin, interval, Subscription } from 'rxjs';

delete (L.Icon.Default.prototype as any)._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
  iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
  shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
});

@Component({
  selector: 'app-tour-execution',
  templateUrl: './tour-execution.component.html',
  styleUrls: ['./tour-execution.component.css'],
})
export class TourExecutionComponent implements OnInit, AfterViewInit, OnDestroy {


  private map!: L.Map;
  private marker?: L.Marker;
  private positionCheckSub?: Subscription;

  tourExecution?: TourExecution;
  keyPoints: KeyPoint[] = [];
  tour?: Tour;
  completedKeyPoints: CompletedKeyPoint[] = [];

  constructor(
    private route: ActivatedRoute,
    private tourExecutionService: TourExecutionService,
    private tourKeyPointService: KeyPointService,
    private tourService: TourService
  ) {}

  ngOnInit(): void {
  const id = this.route.snapshot.paramMap.get('id');
  if (!id) return;

  // 1. Dobavi TourExecution
  this.tourExecutionService.getTourExecution(id).subscribe({
    next: (execution: TourExecution) => {
      this.tourExecution = execution;
      console.log('Loaded tour execution:', execution);

      // 2. Dobavi odgovarajuću turu iz svih tura
      this.tourService.getAllTours().subscribe({
        next: (tours) => {
          const tour = tours.find(t => t.id === execution.tourId);
          if (tour) {
            console.log('Loaded corresponding tour:', tour);
           this.tour = tour;
          } else {
            console.warn('Tour not found for this execution');
          }
        },
        error: err => console.error('Error fetching all tours:', err)
      });

    this.loadKeyPointsAndCompleted(execution.tourId.toString(), id);

    },
    error: err => console.error('Error fetching tour execution:', err)
  });
}

  ngOnDestroy(): void {
    this.positionCheckSub?.unsubscribe(); 
  }

  private loadKeyPointsAndCompleted(tourId: string, executionId: string): void {
  forkJoin({
    keyPoints: this.tourKeyPointService.getKeyPointsByTour(tourId),
    completed: this.tourExecutionService.getCompletedKeyPoints(executionId)
  }).subscribe({
    next: ({ keyPoints, completed }) => {
      this.keyPoints = keyPoints;
      this.completedKeyPoints = completed.map(cp => {
        const kp = keyPoints.find(k => k.id === cp.keyPointId);
        return {
          ...cp,
          keyPointName: kp ? kp.name : 'Unknown'
        };
      });

      this.addKeyPointMarkers(this.keyPoints);

      if (this.tourExecution?.status === 'active') {
        this.startPositionCheck();
      }
    },
    error: err => console.error('Error loading key points or completed KP:', err)
  });
}



  ngAfterViewInit(): void {
    // Restore from localStorage if available
    const storedPos = localStorage.getItem('touristPosition');
    let initialLatLng: [number, number] = [44.8176, 20.4569]; 

    if (storedPos) {
      const { lat, lng } = JSON.parse(storedPos);
      initialLatLng = [lat, lng];
    }

    // Init map
    this.map = L.map('map', {
      center: initialLatLng,
      zoom: 13
    });

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '&copy; OpenStreetMap contributors'
    }).addTo(this.map);

    // Draw marker if saved position exists
    if (storedPos) {
      this.setMarker(initialLatLng[0], initialLatLng[1]);
    }

    // Click to set position
    this.map.on('click', (e: L.LeafletMouseEvent) => {
      this.setMarker(e.latlng.lat, e.latlng.lng);

      // Save in localStorage
      localStorage.setItem('touristPosition', JSON.stringify({
        lat: e.latlng.lat,
        lng: e.latlng.lng
      }));

      console.log('Tourist position saved:', e.latlng);
    });
  }

  private setMarker(lat: number, lng: number) {
    if (this.marker) {
      this.marker.setLatLng([lat, lng]);
    } else {
      this.marker = L.marker([lat, lng]).addTo(this.map);
    }
  }

  private addKeyPointMarkers(keyPoints: KeyPoint[]) {
    if (!this.map || !keyPoints || keyPoints.length === 0) return;

    const waypoints: L.LatLng[] = keyPoints.map(kp => {
      L.marker([kp.latitude, kp.longitude], {
        icon: L.icon({
          iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
          iconSize: [25, 41],
          iconAnchor: [12, 41]
        })
      })
      .addTo(this.map)
      .bindTooltip(`<b>${kp.name}</b><br>${kp.description}`, { permanent: false, direction: 'top' });
      return L.latLng(kp.latitude, kp.longitude);
    });

    // Ruta po ulicama ako ima 2+ KT
    if (waypoints.length >= 2) {
      L.Routing.control({
        waypoints: waypoints,
        router: L.Routing.mapbox(
          'pk.eyJ1IjoidmVsam9vMDIiLCJhIjoiY20yaGV5OHU4MDFvZjJrc2Q4aGFzMTduNyJ9.vSQUDO5R83hcw1hj70C-RA',
          { profile: 'mapbox/walking' }
        ),
        lineOptions: { styles: [{ color: '#1E90FF', weight: 5, opacity: 0.8 }] } as any,
        routeWhileDragging: false,
        showAlternatives: false
      }).addTo(this.map);
    }
  }

  completeTour() {
  throw new Error('Method not implemented.');
  }
  abandonTour() {
  throw new Error('Method not implemented.');
  }

  private startPositionCheck() {
  if (!this.keyPoints || this.keyPoints.length === 0) return;

  this.positionCheckSub = interval(10000).subscribe(() => {
    // Uzmi poslednju poziciju iz localStorage
    const storedPos = localStorage.getItem('touristPosition');
    if (!storedPos) return;

    const { lat, lng } = JSON.parse(storedPos);

    // Filtriraj KT koje nisu još kompletirane
    const incompleteKeyPoints = this.keyPoints.filter(kp =>
      !this.completedKeyPoints.some(ckp => ckp.keyPointId === kp.id)
    );

    if (this.tour?.id) {
      this.tourKeyPointService.getClosestKeyPoint(this.tour.id.toString(), lat, lng).subscribe({
        next: (response) => {
          if (response.keyPoint) {
            console.log('Najbliža ključna tačka:', response.keyPoint);

            // Proveri da li je pronađena tačka već kompletirana
            const alreadyCompleted = this.completedKeyPoints.some(
              ckp => ckp.keyPointId === response.keyPoint!.id 
            );

            if (!alreadyCompleted && this.tourExecution?.id && response.keyPoint.id !== undefined) {
              this.tourExecutionService.completeKeyPoint(
                this.tourExecution.id.toString(),
                response.keyPoint.id.toString()
              ).subscribe({
                next: (completedKP) => {
                  console.log('KeyPoint completed:', completedKP);
                  if(this.tour?.id !== undefined && this.tourExecution?.id !== undefined)
                  {
                       this.loadKeyPointsAndCompleted(this.tour.id.toString(), this.tourExecution.id.toString());
                  }
                 
                },
                error: (err) => console.error('Greška pri kompletiranju tačke:', err)
              });
            }
          } else {
            console.log('Nema ključnih tačaka u blizini');
          }
        },
        error: (err) => console.error('Greška pri proveri najbliže ključne tačke:', err)
      });
    }

    if(this.tourExecution?.id !== undefined)
    {
        this.tourExecutionService.updateActivity(this.tourExecution.id!).subscribe({
    next: (updatedExecution) => {
      console.log('TourExecution updated:', updatedExecution);
      if (this.tourExecution) {
      this.tourExecution.lastActivity = new Date();
    }
    },
    error: (err) => console.error('Error updating activity:', err)
  });
    }

  });
}



  
}
