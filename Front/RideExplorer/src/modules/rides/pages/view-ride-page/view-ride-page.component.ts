import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UtilService } from 'src/modules/shared/services/util.service';
import { DriveDTO } from '../../models/drive-dto';
import { RideService } from '../../services/ride.service';

@Component({
  selector: 'app-view-ride-page',
  templateUrl: './view-ride-page.component.html',
  styleUrls: ['./view-ride-page.component.scss']
})
export class ViewRidePageComponent implements OnInit {

  ride: DriveDTO;

  constructor(private rideService: RideService, 
    private utilService: UtilService,
    private route: ActivatedRoute,
    private snackBarService: SnackBarService) {
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

    const id = parseInt(idStr);
    this.rideService.getRide(id).subscribe((response) => {
      if(response.body)
        this.ride = response.body;
    }, 
    (error) => {
      this.snackBarService.openSnackBar("An error ocured while loading ride");
      this.utilService.navigateToMyProfile();
    });
  }

}
