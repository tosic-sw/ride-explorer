import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UtilService } from 'src/modules/shared/services/util.service';
import { DriveDTO, SearchDTO } from '../../models/drive-dto';
import { RideService } from '../../services/ride.service';

@Component({
  selector: 'app-search-ride-page',
  templateUrl: './search-ride-page.component.html',
  styleUrls: ['./search-ride-page.component.scss']
})
export class SearchRidePageComponent implements OnInit {

  rides: DriveDTO[];
  pageSize: number;
  currentPage: number;
  totalSize: number;
  departureSearch: string;
  destinationSearch: string;
  searched: boolean;

  @ViewChild(PaginationComponent) pagination!: PaginationComponent;

  constructor(
    private router: Router,
    private rideSerive: RideService,
    private snackBarService: SnackBarService
    ) { 
      this.rides = [];
      this.pageSize = 6;
      this.currentPage = 1;
      this.totalSize = 1;
      this.departureSearch = "";
      this.destinationSearch = "";
      this.searched = false;
    }

  ngOnInit(): void {}

  changePage(newPage: any) {
    const newPageNumber = newPage as number;

    const searchDTO: SearchDTO = {
      departure_location: this.departureSearch,
      destination: this.destinationSearch,
      page: newPageNumber - 1,
      size: this.pageSize
    }

    this.rideSerive.searchRides(searchDTO).subscribe((response) => {
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
        this.snackBarService.openSnackBar("An unknown error ocured while searching for rides");
    });
  }

  viewRide(id: number) {
    this.router.navigate([`ridexplorer/rides/view/${id}`])
  }

  reservePlace(id: number) {
    console.log("Should reserve place for passenger");
  }

  searchRide(searchDTO: SearchDTO) {
    this.departureSearch = searchDTO.departure_location;
    this.destinationSearch = searchDTO.destination;

    this.searched = true;

    this.changePage(1);
  }

}
