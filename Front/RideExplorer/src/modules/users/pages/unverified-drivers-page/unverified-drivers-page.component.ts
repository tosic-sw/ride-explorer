import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UserDTO } from '../../models/user-dto';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-unverified-drivers-page',
  templateUrl: './unverified-drivers-page.component.html',
  styleUrls: ['./unverified-drivers-page.component.scss']
})
export class UnverifiedDriversPageComponent implements OnInit {

  users: UserDTO[];
  pageSize: number;
  currentPage: number;
  totalSize: number;
  searchText: string;

  @ViewChild(PaginationComponent) pagination!: PaginationComponent;

  constructor(private userService: UserService, private snackBarService: SnackBarService, private router: Router) { 
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

    this.userService.search(this.searchText, newPageNumber - 1, this.pageSize, "driver", false).subscribe((response: any) => {
      this.users = response.body;
      this.totalSize = Number(response.headers.get("total-elements"));

      if(newPage === 1)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading unverified drivers");
    });
  }

  search(event: any) {
    let text = event as string;
    if(!text) text = "";

    this.searchText = text;

    this.changePage(1);
  }

  examineUser(username: string): void {
    this.router.navigate(["ridexplorer/auth/driver-verification/" + username]);
  }

}
