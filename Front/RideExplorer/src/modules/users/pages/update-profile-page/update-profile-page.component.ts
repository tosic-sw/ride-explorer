import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UtilService } from 'src/modules/shared/services/util.service';
import { UserForUpdateDTO } from '../../models/user-dto';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-update-profile-page',
  templateUrl: './update-profile-page.component.html',
  styleUrls: ['./update-profile-page.component.scss']
})
export class UpdateProfilePageComponent implements OnInit {

  form: FormGroup;
  username: string;
  role: string;

  constructor(
    private fb: FormBuilder,
    private userService: UserService,
    private utilService: UtilService,
    private snackBarService: SnackBarService
  ) {
    this.username = "";
    this.role = "";

    this.form = this.fb.group({
      password: [null, Validators.required],
      firstname: [null, Validators.required],
      lastname: [null, Validators.required],
      email: [null, Validators.required],
      phoneNumber: [null, Validators.required],
    });
  }

  submit() {
    const dto: UserForUpdateDTO = {
      password: this.form.value.password,
      firstname: this.form.value.firstname,
      lastname: this.form.value.lastname,
      email: this.form.value.email,
      phoneNumber: this.form.value.phoneNumber
    };

    this.userService.updateUser(dto).subscribe((response) => {
      if(response.body) {
        let msgReponse: MessageResponse = response.body
        this.snackBarService.openSnackBar(msgReponse.message);
      }
    }, 
    (error) => {
        this.snackBarService.openSnackBar(error.error.message);
    });
  }

  ngOnInit(): void {
    this.userService.getUserForUpdate().subscribe((response) => {
      if(response.body) {
        const userForUpdate: UserForUpdateDTO = response.body;
        this.form.setValue(userForUpdate);
      }
    },
    (error) => {
      const msgReponse: MessageResponse = error.error;
      this.snackBarService.openSnackBar(msgReponse.message);
    });

  }

}
