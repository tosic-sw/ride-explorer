import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UtilService } from 'src/modules/shared/services/util.service';
import { DriveDTO } from '../../models/drive-dto';
import { RideService } from '../../services/ride.service';

@Component({
  selector: 'app-driver-finished-page',
  templateUrl: './driver-finished-page.component.html',
  styleUrls: ['./driver-finished-page.component.scss']
})
export class DriverFinishedPageComponent implements OnInit {

  rides: DriveDTO[];
  pageSize: number;
  currentPage: number;
  totalSize: number;

  @ViewChild(PaginationComponent) pagination!: PaginationComponent;

  constructor(
    private router: Router,
    private utilService: UtilService,
    private rideSerive: RideService,
    private snackBarService: SnackBarService
    ) { 
      this.rides = [];
      this.pageSize = 8;
      this.currentPage = 1;
      this.totalSize = 1;
    }

  ngOnInit(): void {
    this.changePage(this.currentPage);
  }

  changePage(newPage: any) {
    const newPageNumber = newPage as number;
    const username = this.utilService.getLoggedUserUsername();

    this.rideSerive.getFinishedRidesDriver(username, newPageNumber - 1, this.pageSize).subscribe((response) => {
      if(response.body) {
        const drives = response.body;
        this.rides = drives.drives
        this.totalSize = drives.total_elements;
      }
      
      if(newPage === 1 && this.pagination)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading rides");
    });
  }

  viewRide(id: number) {
    this.router.navigate([`ridexplorer/rides/view/${id}`])
  }

  deleteRide(id: number) {
    const username = this.utilService.getLoggedUserUsername();

    this.rideSerive.deleteDrive(username, id).subscribe((response) => {
      this.snackBarService.openSnackBar("Ride successfully deleted");
      this.changePage(1);
    }, 
    (error) => {
      this.snackBarService.openSnackBar("An error ocured while deleting ride");
    })

  }

}
