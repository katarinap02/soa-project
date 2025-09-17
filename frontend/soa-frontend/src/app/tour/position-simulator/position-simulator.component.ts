import { Component, AfterViewInit } from '@angular/core';
import * as L from 'leaflet';

@Component({
  selector: 'app-position-simulator',
  templateUrl: './position-simulator.component.html',
  styleUrls: ['./position-simulator.component.css']
})
export class PositionSimulatorComponent implements AfterViewInit {

  private map!: L.Map;
  private marker?: L.Marker;

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
}
