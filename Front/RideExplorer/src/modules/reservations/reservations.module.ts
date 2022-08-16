import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { VerifiedReservationsPageComponent } from './pages/verified-reservations-page/verified-reservations-page.component';
import { UnverifiedReservationsPageComponent } from './pages/unverified-reservations-page/unverified-reservations-page.component';
import { ReservationsTemplateComponent } from './components/reservations-template/reservations-template.component';
import { ReservationsRoutes } from './reservations.routes';
import { ReactiveFormsModule } from '@angular/forms';
import { SharedModule } from '../shared/shared.module';
import { RouterModule } from '@angular/router';



@NgModule({
  declarations: [
    VerifiedReservationsPageComponent,
    UnverifiedReservationsPageComponent,
    ReservationsTemplateComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(ReservationsRoutes),
    SharedModule
  ]
})
export class ReservationsModule { }
