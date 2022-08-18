import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ReservationDTO } from '../../models/reservation-dtos';

@Component({
  selector: 'app-reservations-passenger-table',
  templateUrl: './reservations-passenger-table.component.html',
  styleUrls: ['./reservations-passenger-table.component.scss']
})
export class ReservationsPassengerTableComponent implements OnInit {

  @Input()
  reservations: ReservationDTO[];

  @Input()
  rideFinished: boolean;

  @Output()
  viewPassengerEvent = new EventEmitter<string>();

  @Output()
  ratePassengerEvent = new EventEmitter<string>();

  @Output()
  complainPassengerEvent = new EventEmitter<string>();

  constructor() { 
    this.reservations = [];
    this.rideFinished = false;
  }

  viewPassenger(username: string) {
    this.viewPassengerEvent.emit(username);
  }

  ratePassenger(username: string) {
    this.ratePassengerEvent.emit(username);
  }

  complainPassenger(username: string) {
    this.complainPassengerEvent.emit(username);
  }

  ngOnInit(): void {}

}
