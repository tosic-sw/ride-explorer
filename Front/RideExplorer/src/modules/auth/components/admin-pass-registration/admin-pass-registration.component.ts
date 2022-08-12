import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { RegistrationDTO } from '../../models/registration-dto';
import { AuthService } from '../../services/auth-service/auth.service';

@Component({
  selector: 'app-admin-pass-registration',
  templateUrl: './admin-pass-registration.component.html',
  styleUrls: ['./admin-pass-registration.component.scss']
})
export class AdminPassRegistrationComponent implements OnInit {

  @Input()
  role: string;

  form: FormGroup;

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router,
    private snackBarService: SnackBarService
  ) {
    this.role = "";
    this.form = this.fb.group({
      username: [null, Validators.required],
      password: [null, Validators.required],
      firstname: [null, Validators.required],
      lastname: [null, Validators.required],
      email: [null, Validators.required],
      phoneNumber: [null, Validators.required],
    });
  }

  submit() {
    const dto: RegistrationDTO = {
      username: this.form.value.username,
      password: this.form.value.password,
      firstname: this.form.value.firstname,
      lastname: this.form.value.lastname,
      email: this.form.value.email,
      phoneNumber: this.form.value.phoneNumber
    };

    this.authService.adminPassRegistration(dto, this.role).subscribe((response) => {
      let msgReponse: MessageResponse;
      if(response.body) {
        msgReponse = response.body;
        this.snackBarService.openSnackBar(msgReponse.message);

        if(this.role === "admin") 
          this.router.navigate(["ridexplorer"]);

        else if (this.role === "passenger") 
          this.router.navigate(["ridexplorer/auth/login"]);
      } 
    }, 
    (error) => {
        this.snackBarService.openSnackBar(error.error.message);
    });
  }

  ngOnInit(): void {}

}
