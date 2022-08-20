import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { ReservationDTO } from 'src/modules/shared/models/reservation-dtos';
import { ReservationService } from 'src/modules/shared/services/reservation.service';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UtilService } from 'src/modules/shared/services/util.service';
import { DriveDTO, UpdateDriveDTO } from '../../models/drive-dto';
import { RideService } from '../../services/ride.service';

@Component({
  selector: 'app-update-ride-page',
  templateUrl: './update-ride-page.component.html',
  styleUrls: ['./update-ride-page.component.scss']
})
export class UpdateRidePageComponent implements OnInit {

  ride: DriveDTO;
  rideId: number;
  username: string;

  reservations: ReservationDTO[];
  pageSize: number;
  currentPage: number;
  totalSize: number;

  @ViewChild(PaginationComponent) pagination!: PaginationComponent;

  constructor(private rideService: RideService, 
    private utilService: UtilService,
    private route: ActivatedRoute,
    private snackBarService: SnackBarService,
    private router: Router,
    private reservationService: ReservationService) {
    this.ride = {
      id: -1,
      driver_username: "",
      departure_location: "",
      destination: "",
      departure_date_time: 0,
      departure_address: "",
      free_places: -1,
      planned_arrival_time: 0,
      note: "",
      finished: false,
      distance: 0,
    }

    this.rideId = -1;
    this.username = "";

    this.reservations = [];
    this.pageSize = 8;
    this.currentPage = 1;
    this.totalSize = 1;
  }

  ngOnInit(): void {
    this.username = this.utilService.getLoggedUserUsername();
    this.rideId = this.getIdFromRoute();
    this.loadRideData();
    this.changePage(this.currentPage);
  }

  private loadRideData(): void {
    this.rideService.getRideUnfinishedDriver(this.rideId, this.username).subscribe((response) => {
      if(response.body)
        this.ride = response.body;
    }, 
    (error) => {
      this.snackBarService.openSnackBar("Ride is finished, or you are not the driver");
      this.utilService.navigateToMyProfile();
    });
  }

  changePage(newPage: any) {
    const newPageNumber = newPage as number;

    this.reservationService.getUnverifiedForDriveAndDriver(this.rideId, newPageNumber - 1, this.pageSize).subscribe((response) => {
      if(response.body) {
        this.reservations = response.body;
        this.totalSize = Number(response.headers.get("total-elements"));
      }

      if(newPage === 1 && this.pagination)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading reservations");
    });
  }

  updateRide(dto: UpdateDriveDTO): void {
    if(dto.departure_address === this.ride.departure_address &&
       dto.free_places === this.ride.free_places &&
       dto.note === this.ride.note) {
        this.snackBarService.openSnackBar("Ride did not change, all fields are same");
        return;
    }

    this.rideService.updateRide(dto).subscribe((response) => {
      this.snackBarService.openSnackBar("Ride successfully updated");
      // Send request to notify that ride has been changed on reseration service
      this.reservationService.notifyDriveChanged(this.ride.id).subscribe((response) => {
        console.log("Notifications for users that have reserations od this ride successfully sent..")
      },
      (error) => {
        console.log("Notifications for users that have reserations od this ride not sent..")
        console.log(error);
      })
    }, 
    (error) => {
      this.snackBarService.openSnackBar("An error ocured while updating ride");
    })
  }

  finishRide(id: any) {
    const username = this.utilService.getLoggedUserUsername();
    this.rideService.finishRide(username, this.ride.id).subscribe((response) => {
      this.snackBarService.openSnackBar("Ride successfully finished");
      this.router.navigate(["/ridexplorer/rides/finished"]);
    },  
    (error) => {
      this.snackBarService.openSnackBar("An error ocured while finishing ride");
    })
  }

  verifyReservation(id: number) {
    console.log()
    this.reservationService.verifyReservation(id).subscribe((response) => {
      let msg = response as unknown as MessageResponse;
      this.snackBarService.openSnackBar(msg.message);
      this.ride.free_places = this.ride.free_places - 1;
      this.changePage(1);
    },
    (error) => {
      const msg: MessageResponse = error.error;
      this.snackBarService.openSnackBar(msg.message);
    })
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

  private getIdFromRoute(): number {
    const idStr: string | null = this.route.snapshot.paramMap.get("id");
    if(!idStr) {
      this.snackBarService.openSnackBar("Invalid ride id");
      this.utilService.navigateToMyProfile();
      return -1;
    }

    if(!this.utilService.isNumber(idStr)) {
      this.snackBarService.openSnackBar("Invalid ride id");
      this.utilService.navigateToMyProfile();
    }
    return parseInt(idStr);
  }

}
