import { Component, OnInit } from '@angular/core';
import { DriveDTO } from '../../models/drive-dto';
import { RideService } from '../../services/ride.service';

@Component({
  selector: 'app-update-ride-page',
  templateUrl: './update-ride-page.component.html',
  styleUrls: ['./update-ride-page.component.scss']
})
export class UpdateRidePageComponent implements OnInit {

  ride: DriveDTO;

  constructor(private rideService: RideService) {
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

  ngOnInit(): void {}

}
