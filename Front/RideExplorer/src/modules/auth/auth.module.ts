import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthRoutes } from './auth.routes';
import { LoginComponent } from './pages/login/login.component';
import { ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { PassengerRegistrationPageComponent } from './pages/passenger-registration-page/passenger-registration-page.component';
import { SharedModule } from '../shared/shared.module';
import { AdminPassRegistrationComponent } from './components/admin-pass-registration/admin-pass-registration.component';
import { AdminRegistrationPageComponent } from './pages/admin-registration-page/admin-registration-page.component';
import { DriverRegistrationPageComponent } from './pages/driver-registration-page/driver-registration-page.component';

@NgModule({
  declarations: [
    LoginComponent, 
    PassengerRegistrationPageComponent,
    AdminPassRegistrationComponent,
    AdminRegistrationPageComponent,
    DriverRegistrationPageComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(AuthRoutes),
    SharedModule
  ],
  providers: []
})
export class AuthModule { }
