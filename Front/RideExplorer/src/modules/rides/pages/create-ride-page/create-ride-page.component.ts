import { Component, OnInit } from '@angular/core';
import { DriveDTO, NewDriveDTO } from '../../models/drive-dto';
import { RideService } from '../../services/ride.service';

@Component({
  selector: 'app-create-ride-page',
  templateUrl: './create-ride-page.component.html',
  styleUrls: ['./create-ride-page.component.scss']
})
export class CreateRidePageComponent implements OnInit {

  constructor(private rideService: RideService) {}

  ngOnInit(): void {}

  createRide(dto: NewDriveDTO) {
    console.log(dto);
  }
}
