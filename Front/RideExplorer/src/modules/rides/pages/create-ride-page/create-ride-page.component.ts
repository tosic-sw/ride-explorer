import { Component, OnInit } from '@angular/core';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UtilService } from 'src/modules/shared/services/util.service';
import { DriveDTO, NewDriveDTO } from '../../models/drive-dto';
import { RideService } from '../../services/ride.service';

@Component({
  selector: 'app-create-ride-page',
  templateUrl: './create-ride-page.component.html',
  styleUrls: ['./create-ride-page.component.scss']
})
export class CreateRidePageComponent implements OnInit {

  constructor(
    private rideService: RideService, 
    private utilService: UtilService,
    private snackBarService: SnackBarService) {}

  ngOnInit(): void {}

  createRide(dto: NewDriveDTO) {
    this.rideService.createRide(dto).subscribe((response) => {
      this.snackBarService.openSnackBar("Ride successfully created");
      this.utilService.navigateToMyProfile();
    },
    (error) => {
      this.snackBarService.openSnackBar("An error ocure while creating ride");
    })
  }
}
