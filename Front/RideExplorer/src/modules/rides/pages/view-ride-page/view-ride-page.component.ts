import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { ComplaintDTO, CreateComplaintDTO } from 'src/modules/shared/models/complaint-dto';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { RatingDTO } from 'src/modules/shared/models/rating-dto';
import { ReservationDTO } from 'src/modules/shared/models/reservation-dtos';
import { ComplaintService } from 'src/modules/shared/services/complaint.service';
import { RatingService } from 'src/modules/shared/services/rating.service';
import { ReservationService } from 'src/modules/shared/services/reservation.service';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UtilService } from 'src/modules/shared/services/util.service';
import { ComplainComponent } from '../../components/complain/complain.component';
import { RateComponent } from '../../components/rate/rate.component';
import { DriveDTO } from '../../models/drive-dto';
import { RideService } from '../../services/ride.service';

@Component({
  selector: 'app-view-ride-page',
  templateUrl: './view-ride-page.component.html',
  styleUrls: ['./view-ride-page.component.scss']
})
export class ViewRidePageComponent implements OnInit {

  ride: DriveDTO;
  rideId: number;
  username: string;

  reservations: ReservationDTO[];
  pageSize: number;
  currentPage: number;
  totalSize: number;

  constructor(private rideService: RideService, 
    private utilService: UtilService,
    private route: ActivatedRoute,
    private snackBarService: SnackBarService,
    private reservationService: ReservationService,
    public dialog: MatDialog,
    private ratingService: RatingService,
    private complaintService: ComplaintService,
    private router: Router) {
      this.rideId = -1;
      this.username = "";
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
    this.rideService.getRide(this.rideId).subscribe((response) => {
      if(response.body)
        this.ride = response.body;
    }, 
    (error) => {
      this.snackBarService.openSnackBar("An error ocured while loading ride");
      this.utilService.navigateToMyProfile();
    });
  }

  changePage(newPage: any) {
    const newPageNumber = newPage as number;

    this.reservationService.getVerifiedForRide(this.rideId, newPageNumber - 1, this.pageSize).subscribe((response) => {
      if(response.body) {
        this.reservations = response.body;
        this.totalSize = Number(response.headers.get("total-elements"));
      }
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading reservations");
    });
  }

  viewPassenger(username: string) {
    this.router.navigate([`/ridexplorer/users/passenger/${username}`]);
  }

  ratePassenger(username: string) {
    if(username === this.utilService.getLoggedUserUsername()) {
      this.snackBarService.openSnackBar("You cannot rate yourself")
      return;
    }
    this.openDialogRate(username);
  }

  complainPassenger(username: string) {
    if(username === this.utilService.getLoggedUserUsername()) {
      this.snackBarService.openSnackBar("You cannot comlpain about yourself")
      return;
    }
    this.openDialogComplain(username);
  }

  openDialogRate(username: string): void {
    const dialogRef = this.dialog.open(RateComponent, {
      width: '550px',
      height: '450px',
      data: username
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);

      const dto: RatingDTO = {
        driveId: this.ride.id,
        positive: result.positive,
        text: result.text,
        evaluated: username
      };

      this.ratingService.createRating(dto).subscribe((response) => {
        if(response.body) {
          const msg: MessageResponse = response.body;
          this.snackBarService.openSnackBar(msg.message);
        }        
      }, 
      (error) => {
        const msg: MessageResponse = error.error;
        this.snackBarService.openSnackBar(msg.message);
      });

    });
  }

  openDialogComplain(username: string): void {
    const dialogRef = this.dialog.open(ComplainComponent, {
      width: '550px',
      height: '420px',
      data: username
    });

    dialogRef.afterClosed().subscribe(result => {

      const dto: CreateComplaintDTO = {
        driveId: this.ride.id,
        accused: username,
        text: result
      };

      this.complaintService.createComplaint(dto).subscribe((response) => {
        if(response.body) {
          const msg: MessageResponse = response.body;
          this.snackBarService.openSnackBar(msg.message);
        }  
      },
      (error) => {
        const msg: MessageResponse = error.error;
        this.snackBarService.openSnackBar(msg.message);
      });
      
    });
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
