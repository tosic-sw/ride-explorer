import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SharedModule } from '../shared/shared.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { UsersRoutes } from './users.routes';
import { UnverifiedDriversPageComponent } from './pages/unverified-drivers-page/unverified-drivers-page.component';
import { DriversPageComponent } from './pages/drivers-page/drivers-page.component';
import { AdminsPageComponent } from './pages/admins-page/admins-page.component';
import { PassengersPageComponent } from './pages/passengers-page/passengers-page.component';
import { UsersTableComponent } from './components/users-table/users-table.component';
import { MatIconModule } from '@angular/material/icon';
import { UnverifiedDriverTableComponent } from './components/unverified-driver-table/unverified-driver-table.component';
import { CarViewerComponent } from './components/car-viewer/car-viewer.component';
import { UserViewerComponent } from './components/user-viewer/user-viewer.component';
import { PassengerPageComponent } from './pages/passenger-page/passenger-page.component';
import { DriverPageComponent } from './pages/driver-page/driver-page.component';
import { RatingViewerComponent } from './components/rating-viewer/rating-viewer.component';
import { UpdateProfilePageComponent } from './pages/update-profile-page/update-profile-page.component';
import { MatDialogModule } from '@angular/material/dialog';
import { MatTooltipModule } from '@angular/material/tooltip';

@NgModule({
  declarations: [
    UnverifiedDriversPageComponent,
    DriversPageComponent,
    AdminsPageComponent,
    PassengersPageComponent,
    UsersTableComponent,
    UnverifiedDriverTableComponent,
    CarViewerComponent,
    UserViewerComponent,
    PassengerPageComponent,
    DriverPageComponent,
    RatingViewerComponent,
    UpdateProfilePageComponent,
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(UsersRoutes),
    SharedModule,
    FormsModule,
    MatIconModule,
    MatDialogModule,
    MatTooltipModule
  ]
})
export class UsersModule { }
