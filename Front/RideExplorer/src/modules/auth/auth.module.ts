import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthRoutes } from './auth.routes';
import { LoginComponent } from './pages/login/login.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { PassengerRegistrationPageComponent } from './pages/passenger-registration-page/passenger-registration-page.component';
import { SharedModule } from '../shared/shared.module';
import { AdminPassRegistrationComponent } from './components/admin-pass-registration/admin-pass-registration.component';
import { AdminRegistrationPageComponent } from './pages/admin-registration-page/admin-registration-page.component';
import { DriverRegistrationPageComponent } from './pages/driver-registration-page/driver-registration-page.component';
import { DriverVerificationPageComponent } from './pages/driver-verification-page/driver-verification-page.component';
import { MatTooltipModule } from '@angular/material/tooltip';

@NgModule({
  declarations: [
    LoginComponent, 
    PassengerRegistrationPageComponent,
    AdminPassRegistrationComponent,
    AdminRegistrationPageComponent,
    DriverRegistrationPageComponent,
    DriverVerificationPageComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(AuthRoutes),
    SharedModule,
    FormsModule,
    MatTooltipModule
  ],
  providers: []
})
export class AuthModule { }
