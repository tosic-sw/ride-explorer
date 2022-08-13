import { Component, OnInit } from '@angular/core';
import { UtilService } from 'src/modules/shared/services/util.service';

@Component({
  selector: 'app-header-driver',
  templateUrl: './header-driver.component.html',
  styleUrls: ['./header-driver.component.scss']
})
export class HeaderDriverComponent implements OnInit {

  username: string;

  constructor(private utilService: UtilService) { 
    this.username = "";
  }

  ngOnInit(): void {
    this.username = this.utilService.getLoggedUserUsername();
  }

}
