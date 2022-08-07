import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MessageResponse } from 'src/modules/shared/models/message-response';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { RegistrationDTO } from '../../models/registration-dto';
import { AuthService } from '../../services/auth-service/auth.service';

@Component({
  selector: 'app-passenger-registration-page',
  templateUrl: './passenger-registration-page.component.html',
  styleUrls: ['./passenger-registration-page.component.scss']
})
export class PassengerRegistrationPageComponent implements OnInit {

  constructor() {}

  ngOnInit(): void {}

}
