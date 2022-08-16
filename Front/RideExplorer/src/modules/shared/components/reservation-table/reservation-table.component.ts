import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ReservationDTO } from '../../models/reservation-dtos';

@Component({
  selector: 'app-reservation-table',
  templateUrl: './reservation-table.component.html',
  styleUrls: ['./reservation-table.component.scss']
})
export class ReservationTableComponent implements OnInit {

  @Input()
  reservations: ReservationDTO[];

  @Output()
  viewRideEvent = new EventEmitter<number>();

  @Output()
  viewDriverEvent = new EventEmitter<string>();

  @Output()
  viewPassengerEvent = new EventEmitter<string>();

  constructor() { 
    this.reservations = [];
  }

  viewRide(id: number) {
    this.viewRideEvent.emit(id);
  }

  viewDriver(username: string) {
    this.viewDriverEvent.emit(username);
  }

  viewPassenger(username: string) {
    this.viewPassengerEvent.emit(username);
  }

  ngOnInit(): void {}

}
