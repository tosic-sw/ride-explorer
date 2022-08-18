import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { ViewRatingDTO } from '../../../shared/models/rating-dto';
import { DriverWithCarDTO } from '../../models/user-dto';
import { RatingService } from '../../../shared/services/rating.service';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-driver-page',
  templateUrl: './driver-page.component.html',
  styleUrls: ['./driver-page.component.scss']
})
export class DriverPageComponent implements OnInit {

  user: DriverWithCarDTO;

  ratings: ViewRatingDTO[];
  pageSize: number;
  currentPage: number;
  totalSize: number;

  @ViewChild(PaginationComponent) pagination!: PaginationComponent;

  constructor(
      private userService: UserService, 
      private snackBarService: SnackBarService, 
      private route: ActivatedRoute,
      private router: Router,
      private ratingService: RatingService
    ) {
    this.user = {
      username: "",
      firstname: "",
      lastname: "",
      email: "",
      phoneNumber: "",
      car: {
        brand: "",
        carModel: "",
        plateNumber: "",
        fuelConsumption: 0.0,
        volume: 0.0,
        power: 0.0
      }
    };

    this.pageSize = 3;
    this.currentPage = 1;
    this.totalSize = 1;

    this.ratings = []
   }

  ngOnInit(): void {
    const username = this.route.snapshot.paramMap.get("username");
    if(!username) {
      this.snackBarService.openSnackBar("Error ocured");
      this.router.navigate(["ridexplorer"]);
      return;
    }

    this.loadDriverData(username);
  }

  loadDriverData(username: string): void {
    this.userService.getDriver(username).subscribe((response) => {
      if(response.body) {
        this.user = response.body;
        console.log(this.user);
        
        this.changePage(this.currentPage)
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

  changePage(newPage: any) {
    let newPageNumber = newPage as number;

    this.ratingService.getRatings(this.user.username, newPageNumber - 1, this.pageSize).subscribe((response: any) => {
      if(response.body) {
        this.ratings = response.body;
        this.totalSize = Number(response.headers.get("total-elements"));
      }
      
      if(newPage === 1 && this.pagination)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading ratings for " + this.user.username);
    });
  }

}
