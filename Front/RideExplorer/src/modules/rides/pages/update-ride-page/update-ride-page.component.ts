import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
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

  constructor(private rideService: RideService, 
    private utilService: UtilService,
    private route: ActivatedRoute,
    private snackBarService: SnackBarService,
    private router: Router) {
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
  }

  ngOnInit(): void {
    const idStr: string | null = this.route.snapshot.paramMap.get("id");
    if(!idStr) return;

    if(!this.utilService.isNumber(idStr)) {
      this.snackBarService.openSnackBar("Invalid ride id");
      this.utilService.navigateToMyProfile();
    }
    const username = this.utilService.getLoggedUserUsername();

    const id = parseInt(idStr);
    this.rideService.getRideUnfinishedDriver(id, username).subscribe((response) => {
      if(response.body)
        this.ride = response.body;
    }, 
    (error) => {
      this.snackBarService.openSnackBar("Ride is finished, or you are not the driver");
      this.utilService.navigateToMyProfile();
    });
  }

  updateRide(dto: UpdateDriveDTO): void {
    if(dto.departure_address === this.ride.departure_address &&
       dto.free_places === this.ride.free_places &&
       dto.note === this.ride.note) {
        this.snackBarService.openSnackBar("Ride did not change, all fields are same");
        return;
    }
    // Nastavi sa pozivom na bekend

    this.rideService.updateRide(dto).subscribe((response) => {
      this.snackBarService.openSnackBar("Ride successfully updated");
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

}
