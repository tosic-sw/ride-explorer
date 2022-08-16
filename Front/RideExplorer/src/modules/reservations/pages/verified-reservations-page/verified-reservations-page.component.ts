import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { ReservationDTO } from 'src/modules/shared/models/reservation-dtos';
import { ReservationService } from 'src/modules/shared/services/reservation.service';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';

@Component({
  selector: 'app-verified-reservations-page',
  templateUrl: './verified-reservations-page.component.html',
  styleUrls: ['./verified-reservations-page.component.scss']
})
export class VerifiedReservationsPageComponent implements OnInit {

  constructor() {}

  ngOnInit(): void {}

}
