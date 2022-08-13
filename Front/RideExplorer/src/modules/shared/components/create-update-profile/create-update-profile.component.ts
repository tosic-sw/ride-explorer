import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RegistrationDTO } from 'src/modules/auth/models/registration-dto';

@Component({
  selector: 'app-create-update-profile',
  templateUrl: './create-update-profile.component.html',
  styleUrls: ['./create-update-profile.component.scss']
})
export class CreateUpdateProfileComponent implements OnInit {

  @Input()
  isCreate: boolean;

  @Output()
  createUpdateEvent = new EventEmitter<RegistrationDTO>();

  form: FormGroup;

  constructor(private fb: FormBuilder,) {
    this.isCreate = true;
    this.form = this.fb.group({
      username: [null, Validators.required],
      password: [null, Validators.required],
      firstname: [null, Validators.required],
      lastname: [null, Validators.required],
      email: [null, Validators.required],
      phoneNumber: [null, Validators.required],
    });
  }

  ngOnInit(): void {
    if(!this.isCreate) {
      this.form.removeControl("username");
    }
  }

  submit() {
    const dto: RegistrationDTO = {
      username: this.isCreate ? this.form.value.username : "",
      password: this.form.value.password,
      firstname: this.form.value.firstname,
      lastname: this.form.value.lastname,
      email: this.form.value.email,
      phoneNumber: this.form.value.phoneNumber
    };

    this.createUpdateEvent.emit(dto);
  }
}
