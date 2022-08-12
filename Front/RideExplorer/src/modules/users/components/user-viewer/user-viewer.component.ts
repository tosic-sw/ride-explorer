import { Component, Input, OnInit } from '@angular/core';
import { UserDTO } from '../../models/user-dto';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-user-viewer',
  templateUrl: './user-viewer.component.html',
  styleUrls: ['./user-viewer.component.scss']
})
export class UserViewerComponent implements OnInit {

  @Input()
  user: UserDTO;

  constructor() { 
    this.user = {
      username: "",
      firstname: "",
      lastname: "",
      email: "",
      phoneNumber: "",
    }
  }

  ngOnInit(): void {}

}
