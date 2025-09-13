import { AfterViewInit, Component, OnInit } from '@angular/core';
import {  KeyPointService } from '../service/key-points.service';
import { ActivatedRoute } from '@angular/router';
import * as L from 'leaflet';
delete (L.Icon.Default.prototype as any)._getIconUrl;
import 'leaflet-routing-machine';
import { KeyPoint } from '../model/keyPoint.model';
L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
  iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
  shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
});

@Component({
  selector: 'app-view-tour-map',
  templateUrl: './view-tour-map.component.html',
  styleUrls: ['./view-tour-map.component.css']
})
export class ViewTourMapComponent implements OnInit, AfterViewInit {
keyPoints: KeyPoint[] = [];
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
          this.loadKeyPointsTable();
      }

      editingKeyPoint: KeyPoint | null = null; // currently editing keypoint
editingMarker!: L.Marker; // marker being moved

loadExistingKeyPoints(): void {
  if (!this.tourId) return;

  this.keyPointService.getKeyPointsByTour(this.tourId).subscribe({
    next: (keyPoints: KeyPoint[]) => {
      if (!keyPoints || keyPoints.length === 0) {
        console.warn('No key points to display.');
        return;
      }

      // Add markers
      const waypoints: L.LatLng[] = keyPoints.map(kp => {
        const marker = L.marker([kp.latitude, kp.longitude])
          .addTo(this.map)
          .bindTooltip(
            `<b>${kp.name}</b><br>${kp.description}`,
            { permanent: false, direction: 'top' }
          );

        return L.latLng(kp.latitude, kp.longitude);
      });

      // Draw route only if there are 2 or more points
      if (waypoints.length >= 2) {
        const routeControl = L.Routing.control({
          waypoints: waypoints,
          router: L.Routing.mapbox('pk.eyJ1IjoidmVsam9vMDIiLCJhIjoiY20yaGV5OHU4MDFvZjJrc2Q4aGFzMTduNyJ9.vSQUDO5R83hcw1hj70C-RA', { profile: 'mapbox/walking' }),
          lineOptions: {
            styles: [{ color: '#1E90FF', weight: 5, opacity: 0.9 }]
          } as any,
          routeWhileDragging: false,
          showAlternatives: false
        }).addTo(this.map);

        routeControl.on('routesfound', (e: any) => {
          const summary = e.routes[0].summary;
          console.log(`Total distance: ${summary.totalDistance / 1000} km`);
        });
      }
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
         const lat = e.latlng.lat;
  const lng = e.latlng.lng;

  if (this.editingKeyPoint) {
    // Update coordinates of the key point being edited
    this.editingKeyPoint.latitude = lat;
    this.editingKeyPoint.longitude = lng;

    // Move the editing marker to the new position
    if (this.editingMarker) {
      this.editingMarker.setLatLng(e.latlng);
    }
  } else {
    // Add new key point
    this.keyPoint.latitude = lat;
    this.keyPoint.longitude = lng;

    L.marker([lat, lng]).addTo(this.map);
  }
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
            this.loadExistingKeyPoints();
            this.loadKeyPointsTable();
          },
          error: (err) => {
            console.error(err);
            alert('Error adding key point');
            // U slučaju greške ukloni marker
            this.map.removeLayer(newMarker);
          }
        });
      }

      loadKeyPointsTable(): void {
  if (!this.tourId) return;

  this.keyPointService.getKeyPointsByTour(this.tourId).subscribe({
    next: (points: KeyPoint[]) => {
      this.keyPoints = points;
  
      // optionally refresh markers/routes here
    },
    error: (err) => console.error('Error loading key points', err)
  });
}

deleteKeyPoint(kp: KeyPoint): void {

 
    this.keyPointService.deleteKeyPoint(kp.id!).subscribe({
      next: () => {
        alert('KeyPoint deleted!');
        this.loadExistingKeyPoints();
        this.loadKeyPointsTable(); 
         window.location.reload();
   
      },
      error: (err) => console.error('Delete failed', err)
    });
  }

  get currentKeyPoint(): KeyPoint {
  return this.editingKeyPoint || this.keyPoint;
}

editKeyPoint(kp: KeyPoint) {
  this.editingKeyPoint = { ...kp }; // clone to avoid direct changes
  this.map.setView([kp.latitude, kp.longitude], 15);

  // optionally add a draggable marker for editing
  if (this.editingMarker) {
    this.map.removeLayer(this.editingMarker);
  }

  this.editingMarker = L.marker([kp.latitude, kp.longitude], { draggable: true })
    .addTo(this.map)
    .bindTooltip(`${kp.name}`, { permanent: true })
    .on('dragend', (e: any) => {
      this.editingKeyPoint!.latitude = e.target.getLatLng().lat;
      this.editingKeyPoint!.longitude = e.target.getLatLng().lng;
    });


}

cancelEdit() {
  this.editingKeyPoint = null;
  if (this.editingMarker) {
    this.map.removeLayer(this.editingMarker);
  }
}


updateKeyPoint() {
  if (!this.editingKeyPoint) return;

  this.keyPointService.updateKeyPoint(this.editingKeyPoint.id!, this.editingKeyPoint).subscribe({
    next: () => {
      alert('KeyPoint updated!');
      this.cancelEdit();
      this.loadExistingKeyPoints();
      this.loadKeyPointsTable();
      window.location.reload();
    },
    error: (err) => console.error('Update failed', err)
  });
}



}
