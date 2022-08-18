import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ComplaintDTO } from '../../../shared/models/complaint-dto';

@Component({
  selector: 'app-complaint-card',
  templateUrl: './complaint-card.component.html',
  styleUrls: ['./complaint-card.component.scss']
})
export class ComplaintCardComponent implements OnInit {

  @Input()
  complaint: ComplaintDTO;

  @Output()
  viewProfileEvent = new EventEmitter<string>();

  constructor() {
    this.complaint = {
      id: -1,
      accuser: "",
      accused: "",
      driveId: -1,
      text: "",
      createdAt: -1
    };
   }

  ngOnInit(): void {}

   viewProfile(username: string): void {
    this.viewProfileEvent.emit(username);
   }

}
