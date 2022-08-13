import { Component, OnInit } from '@angular/core';
import { UtilService } from 'src/modules/shared/services/util.service';

@Component({
  selector: 'app-header-passenger',
  templateUrl: './header-passenger.component.html',
  styleUrls: ['./header-passenger.component.scss']
})
export class HeaderPassengerComponent implements OnInit {

  username: string;

  constructor(private utilService: UtilService) { 
    this.username = "";
  }

  ngOnInit(): void {
    this.username = this.utilService.getLoggedUserUsername();
  }

}
