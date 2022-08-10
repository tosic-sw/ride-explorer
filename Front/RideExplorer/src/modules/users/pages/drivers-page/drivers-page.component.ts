import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UserDTO } from '../../models/user-dto';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-drivers-page',
  templateUrl: './drivers-page.component.html',
  styleUrls: ['./drivers-page.component.scss']
})
export class DriversPageComponent implements OnInit {

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

    this.userService.search(this.searchText, newPageNumber - 1, this.pageSize, "driver").subscribe((response: any) => {
      this.users = response.body;
      this.totalSize = Number(response.headers.get("total-elements"));
      
      if(newPage === 1)
        this.pagination.reset();
    },
    (error) => {
      if(error.status === 500)
        this.snackBarService.openSnackBar("An unknown error ocured while loading drivers");
    });
  }

  search(event: any) {
    let text = event as string;
    if(!text) text = "";

    this.searchText = text;

    this.changePage(1);
  }

  banUser(username: string) {
    this.userService.banUser(username, "driver").subscribe((response) => {
      if(response) {
        const idx = this.users.findIndex(user => user.username === username);
        this.users.splice(idx, 1);

        const msg = response as unknown as MessageResponse;
        this.snackBarService.openSnackBar(msg.message);
      }
    },
    (error) => {
      console.log(error);
      if(error.status === 404) {
        this.snackBarService.openSnackBar("Driver with given username not found");
      }
    });
  }

  deleteUser(username: string) {
    this.userService.deleteUser(username, "driver").subscribe((response) => {
      if(response.body) {
        const idx = this.users.findIndex(user => user.username === username);
        this.users.splice(idx, 1);
        
        const msg: MessageResponse = response.body;
        this.snackBarService.openSnackBar(msg.message);
      }
    },
    (error) => {
      console.log(error);
      if(error.status === 404) {
        this.snackBarService.openSnackBar("Driver with given username not found");
      }
    });
  }

  viewProfile(username: string) {
    this.router.navigate(["ridexplorer/users/driver/" + username]);
  }

 }
