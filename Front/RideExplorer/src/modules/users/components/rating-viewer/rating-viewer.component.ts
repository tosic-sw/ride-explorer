import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ViewRatingDTO } from '../../../shared/models/rating-dto';

@Component({
  selector: 'app-rating-viewer',
  templateUrl: './rating-viewer.component.html',
  styleUrls: ['./rating-viewer.component.scss']
})
export class RatingViewerComponent implements OnInit {

  @Input()
  rating: ViewRatingDTO;

  constructor(private router: Router) {
    this.rating = {
      id: -1,
      evaluator: "",
      evaluated: "",
      driveId: -1,
      positive: false,
      text: ""
    };
   }

  ngOnInit(): void {}

   viewProfile(username: string): void {
    console.log("Viewing profile of: " + username);
   }

}
