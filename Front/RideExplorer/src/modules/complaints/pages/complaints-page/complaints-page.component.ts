import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { ComplaintDTO } from '../../models/complaint-dto';
import { ComplaintService } from '../../services/complaint.service';

@Component({
  selector: 'app-complaints-page',
  templateUrl: './complaints-page.component.html',
  styleUrls: ['./complaints-page.component.scss']
})
export class ComplaintsPageComponent implements OnInit {

  complaints: ComplaintDTO[]
  pageSize: number;
  currentPage: number;
  totalSize: number;
  searchText: string;

  @ViewChild(PaginationComponent) pagination!: PaginationComponent;

  constructor(private router: Router, private complaintService: ComplaintService, private snackBarService: SnackBarService) { 
    this.complaints = [];
    this.pageSize = 4;
    this.currentPage = 1;
    this.totalSize = 1;
    this.searchText = "";
  }

  ngOnInit(): void {
    this.changePage(this.currentPage);
  }

  changePage(newPage: any) {
    let newPageNumber = newPage as number;

    this.complaintService.getComplaints(newPageNumber - 1, this.pageSize).subscribe((response: any) => {
      this.complaints = response.body;
      this.totalSize = Number(response.headers.get("total-elements"));
      
      if(newPage === 1)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading complaints");
    });
  }

  viewProfile(event: any) {
    console.log(event);
  }

}
