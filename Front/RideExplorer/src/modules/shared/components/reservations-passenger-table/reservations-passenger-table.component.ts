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

  @Input()
  viewer: string;

  @Output()
  viewPassengerEvent = new EventEmitter<string>();

  @Output()
  ratePassengerEvent = new EventEmitter<string>();

  @Output()
  complainPassengerEvent = new EventEmitter<string>();

  constructor() { 
    this.reservations = [];
    this.rideFinished = false;
    this.viewer = "";
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

  canComplainAndRate(): boolean {
    for(let reservation of this.reservations) {

      if(reservation.passengerUsername === this.viewer)
        return true;
      
      if(reservation.driverUsername === this.viewer)
        return true;
    }
    
    return false;
  }

  ngOnInit(): void {}

}
