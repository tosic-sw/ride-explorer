import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { DriverWithCarDTO } from '../../models/user-dto';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-driver-page',
  templateUrl: './driver-page.component.html',
  styleUrls: ['./driver-page.component.scss']
})
export class DriverPageComponent implements OnInit {

  user: DriverWithCarDTO;

  constructor(
      private userService: UserService, 
      private snackBarService: SnackBarService, 
      private route: ActivatedRoute,
      private router: Router
    ) {
    this.user = {
      username: "",
      firstname: "",
      lastname: "",
      email: "",
      car: {
        brand: "",
        carModel: "",
        plateNumber: "",
        fuelConsumption: 0.0,
        volume: 0.0,
        power: 0.0
      }
    };
   }

  ngOnInit(): void {
    const username = this.route.snapshot.paramMap.get("username");
    if(!username) {
      this.snackBarService.openSnackBar("Error ocured");
      this.router.navigate(["ridexplorer"]);
      return;
    }
    this.userService.getDriver(username).subscribe((response) => {
      if(response.body) {
        this.user = response.body;
        console.log(this.user);
      }
    },
    (error) => {
      if(error.status === 404) {
        this.snackBarService.openSnackBar("Driver not found");
      }
      else {
        this.snackBarService.openSnackBar("Unknown error happend while loading driver")  
      }
      this.router.navigate(["ridexplorer"]);
    });
  }

}
