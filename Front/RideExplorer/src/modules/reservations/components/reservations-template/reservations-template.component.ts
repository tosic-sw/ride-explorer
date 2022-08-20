import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { ReservationDTO } from 'src/modules/shared/models/reservation-dtos';
import { ReservationService } from 'src/modules/shared/services/reservation.service';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';

@Component({
  selector: 'app-reservations-template',
  templateUrl: './reservations-template.component.html',
  styleUrls: ['./reservations-template.component.scss']
})
export class ReservationsTemplateComponent implements OnInit {

  reservations: ReservationDTO[]
  pageSize: number;
  currentPage: number;
  totalSize: number;

  @Input()
  verified: boolean;

  @ViewChild(PaginationComponent) pagination!: PaginationComponent;

  constructor(private router: Router, private reservationService: ReservationService, private snackBarService: SnackBarService) { 
    this.reservations = [];
    this.pageSize = 10;
    this.currentPage = 1;
    this.totalSize = 1;
    this.verified = true;
  }

  ngOnInit(): void {
    this.changePage(this.currentPage);
  }

  changePage(newPage: any) {
    let newPageNumber = newPage as number;

    if(this.verified)
      this.loadVerifiedForPassenger(newPageNumber);
    else
      this.loadUnverifiedForPassenger(newPage) 
  }

  loadVerifiedForPassenger(newPageNumber: number) {
    this.reservationService.getVerifiedForPassenger(newPageNumber - 1, this.pageSize).subscribe((response: any) => {
      if(response.body) {
        this.reservations = response.body;
        this.totalSize = Number(response.headers.get("total-elements"));
      }
      if(newPageNumber === 1 && this.pagination)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading reservations");
    });
  }

  loadUnverifiedForPassenger(newPageNumber: number) {
    this.reservationService.getUnverifiedForPassenger(newPageNumber - 1, this.pageSize).subscribe((response: any) => {
      if(response.body) {
        this.reservations = response.body;
        this.totalSize = Number(response.headers.get("total-elements"));
      }
      if(newPageNumber === 1 && this.pagination)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading reservations");
    });
  }

  viewRide(id: number) {
    this.router.navigate([`/ridexplorer/rides/view/${id}`])
  }

  viewDriver(username: string) {
    this.router.navigate([`/ridexplorer/users/driver/${username}`])
  }

  viewPassenger(username: string) {
    this.router.navigate([`/ridexplorer/users/passenger/${username}`])
  }

  deleteReservation(id: number) {
    this.reservationService.deleteReservation(id).subscribe((response) => {
      if(response.body) {
        const msg: MessageResponse = response.body;
        this.snackBarService.openSnackBar(msg.message);
        this.changePage(1);
      }
    },
    (error) => {  
      console.log(error);
      if(error.error) {
        const msg: MessageResponse = error.error;
        this.snackBarService.openSnackBar(msg.message);
      }
    })
  }

}
