import { Component, OnInit, AfterViewInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import * as L from 'leaflet';
import { KeyPointService } from '../service/key-points.service';
import { KeyPoint } from '../model/keyPoint.model';

delete (L.Icon.Default.prototype as any)._getIconUrl;

L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
  iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
  shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
});

@Component({
  selector: 'app-tour-keypoints',
  templateUrl: './key-points.component.html',
  styleUrls: ['./key-points.component.css']
})
export class KeyPointsComponent implements OnInit, AfterViewInit {

  tourId: string = '';
  keyPoint: KeyPoint = {
    tourId: '',
    name: '',
    description: '',
    latitude: 0,
    longitude: 0,
    imageUrl: ''
  };

  map!: L.Map;

  constructor(private route: ActivatedRoute, private keyPointService: KeyPointService) { }

  ngOnInit(): void {
    this.tourId = this.route.snapshot.paramMap.get('id') || '';
    this.keyPoint.tourId = this.tourId;
  }

  ngAfterViewInit(): void {
    this.initMap();
    this.loadExistingKeyPoints();
  }

//   loadExistingKeyPoints(): void {
//   if (!this.tourId) return;

//   this.keyPointService.getKeyPointsByTour(this.tourId).subscribe({
//     next: (keyPoints: KeyPoint[]) => {
//       keyPoints.forEach(kp => {
//         L.marker([kp.latitude, kp.longitude])
//           .addTo(this.map)
//           .bindPopup(`<b>${kp.name}</b><br>${kp.description}`);
//       });
//     },
//     error: (err) => {
//       console.error('Error loading key points', err);
//     }
//   });
// }

loadExistingKeyPoints(): void {
  if (!this.tourId) return;

  this.keyPointService.getKeyPointsByTour(this.tourId).subscribe({
    next: (keyPoints: KeyPoint[]) => {
      keyPoints.forEach(kp => {
        L.marker([kp.latitude, kp.longitude])
          .addTo(this.map)
          .bindTooltip(
            `<b>${kp.name}</b><br>${kp.description}`, 
            { permanent: false, direction: 'top' } // tooltip iznad markera
          );
      });
    },
    error: (err) => {
      console.error('Error loading key points', err);
    }
  });
}


  initMap(): void {
    this.map = L.map('map').setView([44.8176, 20.4569], 13); // Beograd primer

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: 'Map data © <a href="https://openstreetmap.org">OpenStreetMap</a>'
    }).addTo(this.map);

    this.map.on('click', (e: any) => {
      this.keyPoint.latitude = e.latlng.lat;
      this.keyPoint.longitude = e.latlng.lng;

      // Dodavanje markera
      L.marker([e.latlng.lat, e.latlng.lng]).addTo(this.map);
    });
  }

addKeyPoint(): void {
  if (!this.tourId) return;

  // Marker koji je dodat klikom
  const newMarker = L.marker([this.keyPoint.latitude, this.keyPoint.longitude]).addTo(this.map);

  this.keyPointService.addKeyPoint(this.keyPoint).subscribe({
    next: () => {
      alert('Key point added successfully');
      this.keyPoint = { tourId: this.tourId, name: '', description: '', latitude: 0, longitude: 0, imageUrl: '' };
    },
    error: (err) => {
      console.error(err);
      alert('Error adding key point');
      // U slučaju greške ukloni marker
      this.map.removeLayer(newMarker);
    }
  });
}

}
