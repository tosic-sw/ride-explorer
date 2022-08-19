import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { RatingData } from '../../../rides/models/modal-data';
import { RatingModalData } from '../../models/rating-dto';

@Component({
  selector: 'app-rate',
  templateUrl: './rate.component.html',
  styleUrls: ['./rate.component.scss']
})
export class RateComponent implements OnInit {

  form: FormGroup;

  constructor(
    private dialogRef: MatDialogRef<RateComponent>, 
    @Inject(MAT_DIALOG_DATA) public data: RatingModalData,
    private fb: FormBuilder) { 
    
    this.form = this.fb.group({
      text: [data.text, Validators.required],
      opinion: [data.positive ? "Positive" : "Negative", Validators.required]
    });

  }

  submit() {
    const ratingData: RatingData = {
      positive: this.form.value.opinion === "Positive" ? true : false,
      text: this.form.value.text
    };

    this.dialogRef.close(ratingData);
  }

  ngOnInit(): void {}

}
