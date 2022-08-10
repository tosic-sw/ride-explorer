import { EventEmitter, Component, Input, OnInit, Output } from '@angular/core';
import { UserDTO } from '../../models/user-dto';


@Component({
  selector: 'app-users-table',
  templateUrl: './users-table.component.html',
  styleUrls: ['./users-table.component.scss']
})
export class UsersTableComponent implements OnInit {

  @Input()
  users: UserDTO[];

  @Input()
  role: string;

  @Output()
  banUserEvent = new EventEmitter<string>();
  
  @Output()
  deleteUserEvent = new EventEmitter<string>();

  @Output()
  viewProfileEvent = new EventEmitter<string>();

  constructor() {
    this.users = []
    this.role = "";
  }

  ngOnInit(): void {}

  banUser(username: string): void {
    this.banUserEvent.emit(username);
  }

  deleteUser(username: string): void {
    this.deleteUserEvent.emit(username);
  }

  viewProfile(username: string): void {
    this.viewProfileEvent.emit(username);
  }

}
