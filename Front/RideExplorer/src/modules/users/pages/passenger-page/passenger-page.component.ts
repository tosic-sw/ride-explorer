import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UserDTO } from '../../models/user-dto';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-passenger-page',
  templateUrl: './passenger-page.component.html',
  styleUrls: ['./passenger-page.component.scss']
})
export class PassengerPageComponent implements OnInit {

  user: UserDTO;

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
      email: ""
    };
   }

  ngOnInit(): void {
    const username = this.route.snapshot.paramMap.get("username");
    if(!username) {
      this.snackBarService.openSnackBar("Error ocured");
      this.router.navigate(["ridexplorer"]);
      return;
    }
    this.userService.getPassenger(username).subscribe((response) => {
      if(response.body) {
        this.user = response.body;
        console.log(this.user);
      }
    },
    (error) => {
      if(error.status === 404) {
        this.snackBarService.openSnackBar("Passenger not found");
      }
      else {
        this.snackBarService.openSnackBar("Unknown error happend while loading passenger")  
      }
      this.router.navigate(["ridexplorer"]);
    });
  }

}
