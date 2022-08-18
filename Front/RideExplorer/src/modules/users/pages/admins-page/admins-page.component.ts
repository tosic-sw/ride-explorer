import { Component, OnInit, ViewChild } from '@angular/core';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UserDTO } from '../../models/user-dto';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-admins-page',
  templateUrl: './admins-page.component.html',
  styleUrls: ['./admins-page.component.scss']
})
export class AdminsPageComponent implements OnInit {

  users: UserDTO[];
  pageSize: number;
  currentPage: number;
  totalSize: number;
  searchText: string;

  @ViewChild(PaginationComponent) pagination!: PaginationComponent;

  constructor(private userService: UserService, private snackBarService: SnackBarService) { 
    this.users = [];
    this.pageSize = 8;
    this.currentPage = 1;
    this.totalSize = 1;
    this.searchText = "";
  }

  ngOnInit(): void {
    this.changePage(this.currentPage);
  }

  changePage(newPage: any) {
    let newPageNumber = newPage as number;

    this.userService.search(this.searchText, newPageNumber - 1, this.pageSize, "admin").subscribe((response: any) => {
      if(response.body) {
        this.users = response.body;
        this.totalSize = Number(response.headers.get("total-elements"));
      }

      if(newPage === 1 && this.pagination)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading admins");
    });
  }

  search(event: any) {
    let text = event as string;
    if(!text) text = "";

    this.searchText = text;

    this.changePage(1);
  }

}
