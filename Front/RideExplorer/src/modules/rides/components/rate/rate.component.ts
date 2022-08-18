import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { RatingData } from '../../models/modal-data';

@Component({
  selector: 'app-rate',
  templateUrl: './rate.component.html',
  styleUrls: ['./rate.component.scss']
})
export class RateComponent implements OnInit {

  form: FormGroup;

  constructor(
    private dialogRef: MatDialogRef<RateComponent>, 
    @Inject(MAT_DIALOG_DATA) public data: string,
    private fb: FormBuilder) { 
    
    this.form = this.fb.group({
      text: [null, Validators.required],
      opinion: ["Positive", Validators.required]
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
