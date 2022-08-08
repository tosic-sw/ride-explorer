import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { UserDTO } from '../../models/user-dto';

@Component({
  selector: 'app-unverified-driver-table',
  templateUrl: './unverified-driver-table.component.html',
  styleUrls: ['./unverified-driver-table.component.scss']
})
export class UnverifiedDriverTableComponent implements OnInit {

  @Input()
  users: UserDTO[];
  
  @Output()
  examineUserEvent = new EventEmitter<string>();

  constructor() {
    this.users = []
  }

  ngOnInit(): void {}

  examineUser(username: string): void {
    this.examineUserEvent.emit(username);
  }
  
}
