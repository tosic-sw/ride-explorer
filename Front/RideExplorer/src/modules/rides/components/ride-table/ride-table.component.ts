import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { DriveDTO } from '../../models/drive-dto';

@Component({
  selector: 'app-ride-table',
  templateUrl: './ride-table.component.html',
  styleUrls: ['./ride-table.component.scss']
})
export class RideTableComponent implements OnInit {

  @Input()
  rides: DriveDTO[];

  @Input()
  finished: boolean;

  @Output()
  viewRideEvent = new EventEmitter<number>();

  @Output()
  updateRideEvent = new EventEmitter<number>();

  @Output()
  deleteRideEvent = new EventEmitter<number>();

  constructor() { 
    this.rides = [];
    this.finished = false;
  }
  
  viewRide(id: number) {
    this.viewRideEvent.emit(id);
  }

  updateRide(id: number) {
    this.updateRideEvent.emit(id);
  }

  deleteRide(id: number) {
    this.deleteRideEvent.emit(id);
  }

  ngOnInit(): void {}

}
