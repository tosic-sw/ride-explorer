import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { DriveDTO } from '../../models/drive-dto';

@Component({
  selector: 'app-ride-search-table',
  templateUrl: './ride-search-table.component.html',
  styleUrls: ['./ride-search-table.component.scss']
})
export class RideSearchTableComponent implements OnInit {

  @Input()
  rides: DriveDTO[];

  @Output()
  reservePlaceEvent = new EventEmitter<DriveDTO>();

  @Output()
  viewRideEvent = new EventEmitter<number>();

  constructor() { 
    this.rides = [];
  }
  
  reservePlace(ride: DriveDTO) {
    this.reservePlaceEvent.emit(ride);
  }

  viewRide(id: number) {
    this.viewRideEvent.emit(id);
  }

  ngOnInit(): void {}

}
