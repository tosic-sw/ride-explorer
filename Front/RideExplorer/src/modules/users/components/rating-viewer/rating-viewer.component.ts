import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { RateComponent } from 'src/modules/shared/components/rate/rate.component';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { RatingService } from 'src/modules/shared/services/rating.service';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { RatingDTO, ViewRatingDTO } from '../../../shared/models/rating-dto';

@Component({
  selector: 'app-rating-viewer',
  templateUrl: './rating-viewer.component.html',
  styleUrls: ['./rating-viewer.component.scss']
})
export class RatingViewerComponent implements OnInit {

  @Input()
  rating: ViewRatingDTO;

  @Input()
  viewer: string;

  @Output()
  ratingUpdatedEvent = new EventEmitter<void>()

  constructor(
    public dialog: MatDialog, 
    private ratingService: RatingService, 
    private snackBarService: SnackBarService) {
    this.rating = {
      id: -1,
      evaluator: "",
      evaluated: "",
      driveId: -1,
      positive: false,
      text: ""
    };

    this.viewer = "";
  }

  updateRating(): void {

    const dialogRef = this.dialog.open(RateComponent, {
      width: '550px',
      height: '450px',
      data: {
        text: this.rating.text,
        positive: this.rating.positive
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);

      const dto: RatingDTO = {
        driveId: this.rating.driveId,
        positive: result.positive,
        text: result.text,
        evaluated: this.rating.evaluated
      };

      this.ratingService.updateRating(this.rating.id, dto).subscribe((response) => {
        if(response.body) {
          const msg: MessageResponse = response.body;
          this.snackBarService.openSnackBar(msg.message);

          this.ratingUpdatedEvent.emit();
        }        
      }, 
      (error) => {
        const msg: MessageResponse = error.error;
        this.snackBarService.openSnackBar(msg.message);
      });

    });
  }

  ngOnInit(): void {}



}
