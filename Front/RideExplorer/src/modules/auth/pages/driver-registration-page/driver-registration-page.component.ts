import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { DriverRegistrationDTO, RegistrationDTO } from '../../models/registration-dto';
import { AuthService } from '../../services/auth-service/auth.service';

@Component({
  selector: 'app-driver-registration-page',
  templateUrl: './driver-registration-page.component.html',
  styleUrls: ['./driver-registration-page.component.scss']
})
export class DriverRegistrationPageComponent implements OnInit {
  
  form: FormGroup;

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router,
    private snackBarService: SnackBarService
  ) {
    this.form = this.fb.group({
      username: [null, Validators.required],
      password: [null, Validators.required],
      firstname: [null, Validators.required],
      lastname: [null, Validators.required],
      email: [null, Validators.required],
      plateNumber: [null, Validators.required],
      brand: [null, Validators.required],
      carModel: [null, Validators.required],
      fuelConsumption: ['', [Validators.required, Validators.min(1), Validators.pattern('^(0|[1-9][0-9]*)$')]],
      volume: ['', [Validators.required, Validators.min(1), Validators.pattern('^(0|[1-9][0-9]*)$')]],
      power: ['', [Validators.required, Validators.min(1), Validators.pattern('^(0|[1-9][0-9]*)$')]],
    });
  }

  submit() {
    const dto: DriverRegistrationDTO = {
      username: this.form.value.username,
      password: this.form.value.password,
      firstname: this.form.value.firstname,
      lastname: this.form.value.lastname,
      email: this.form.value.email,
      car: {
        plateNumber: this.form.value.plateNumber,
        brand: this.form.value.brand,
        carModel: this.form.value.carModel,
        fuelConsumption: this.form.value.fuelConsumption,
        volume: this.form.value.volume,
        power: this.form.value.power
      }
    };

    this.authService.driverRegistration(dto).subscribe((response) => {
      let msgReponse: MessageResponse;
      if(response.body) {
        msgReponse = response.body;
        this.snackBarService.openSnackBar(msgReponse.message);
        this.router.navigate(["ridexplorer/auth/login"]);
      } 
    }, 
    (error) => {
        this.snackBarService.openSnackBar(error.error.message);
    });
  }

  ngOnInit(): void {}

}
