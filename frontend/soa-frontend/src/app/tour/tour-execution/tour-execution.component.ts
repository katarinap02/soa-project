import { AfterViewInit, ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import * as L from 'leaflet';
import { TourExecution } from '../model/tourExecution.model';
import { ActivatedRoute } from '@angular/router';
import { TourExecutionService } from '../service/tour-execution.service';
import { KeyPoint } from '../model/keyPoint.model';
import { KeyPointService } from '../service/key-points.service';

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
export class TourExecutionComponent implements OnInit, AfterViewInit {

  private map!: L.Map;
  private marker?: L.Marker;

  tourExecution?: TourExecution;
  keyPoints?: KeyPoint[];

  constructor(
    private route: ActivatedRoute,
    private tourExecutionService: TourExecutionService,
    private tourKeyPointService: KeyPointService
  ) {}

  ngOnInit(): void {
  const id = this.route.snapshot.paramMap.get('id');
  if (id) {
    this.tourExecutionService.getTourExecution(id).subscribe({
      next: (data: TourExecution) => {
        this.tourExecution = data;
        console.log('Loaded tour execution:', data);

        // Kada dobijemo execution, povucemo i keypoints
        this.tourKeyPointService.getKeyPointsByTour(data.tourId.toString())
          .subscribe({
            next: (keyPoints: KeyPoint[]) => {
              console.log('Loaded key points:', keyPoints);
              this.keyPoints = keyPoints;
              this.addKeyPointMarkers(keyPoints);
            },
            error: err => {
              console.error('Error fetching key points:', err);
            }
          });
      },
      error: err => {
        console.error('Error fetching tour execution:', err);
      }
    });
  }
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


  
}
