import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-complain',
  templateUrl: './complain.component.html',
  styleUrls: ['./complain.component.scss']
})
export class ComplainComponent implements OnInit {

  form: FormGroup;

  constructor(
    private dialogRef: MatDialogRef<ComplainComponent>, 
    @Inject(MAT_DIALOG_DATA) public data: string,
    private fb: FormBuilder) { 
    
    this.form = this.fb.group({
      text: [null, Validators.required]
    });

  }

  submit() {
    this.dialogRef.close(this.form.value.text);
  }

  ngOnInit(): void {}
}
