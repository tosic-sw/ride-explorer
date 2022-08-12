import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { DriverRegistrationDTO } from '../../models/registration-dto';
import { AuthService } from '../../services/auth-service/auth.service';
import { ActivatedRoute } from '@angular/router';
import { UserSharedService } from 'src/modules/shared/services/user-shared.service';
import { Message } from '@angular/compiler/src/i18n/i18n_ast';
import { HttpResponse } from '@angular/common/http';
import { DriverWithCarDTO } from 'src/modules/shared/models/driver-shared-dto';

@Component({
  selector: 'app-driver-verification-page',
  templateUrl: './driver-verification-page.component.html',
  styleUrls: ['./driver-verification-page.component.scss']
})
export class DriverVerificationPageComponent implements OnInit {

  username: string = "ciao";

  driverDTO: DriverWithCarDTO;

  constructor(
    private authService: AuthService, 
    private router: Router,
    private snackBarService: SnackBarService,
    private route: ActivatedRoute,
    private userSharedService: UserSharedService
    ) {
      this.driverDTO = {
        username: "",
        password: "",
        firstname: "",
        lastname: "",
        email: "",
        phoneNumber: "",
        car: {
          plateNumber: "",
          brand: "",
          carModel: "",
          fuelConsumption: 0.0,
          volume: 0.0,
          power: 0.0,
        }
      }
    }

  ngOnInit(): void {
    let username = this.route.snapshot.paramMap.get("username");
    if(!username) {
      this.snackBarService.openSnackBar("An error ocured while accesing driver verification page");
      this.router.navigate(["ridexplorer"]);
      return;
    }

    this.userSharedService.getUnverifiedDriver(username).subscribe((response) => {
      if(response.body)
        this.driverDTO = response.body;
    }, 
    (error) => {
      this.snackBarService.openSnackBar(error.error.message);
      console.log(error)
    })
  }

  submit() {
    let username: string;
    if(this.driverDTO)
      username = this.driverDTO.username;
    else 
      return;

    this.authService.driverVerification(username).subscribe((response: HttpResponse<MessageResponse>) => {
      let msg = response as unknown as MessageResponse;
      this.snackBarService.openSnackBar(msg.message);
      this.router.navigate(["ridexplorer"]);
    }, 
    (error) => {
        this.snackBarService.openSnackBar(error.error.message);
    });
  }

  

}
