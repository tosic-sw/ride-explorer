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
import { UnverifiedDriverTableComponent } from './components/unverified-driver-table/unverified-driver-table.component'


@NgModule({
  declarations: [
    UnverifiedDriversPageComponent,
    DriversPageComponent,
    AdminsPageComponent,
    PassengersPageComponent,
    UsersTableComponent,
    UnverifiedDriverTableComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(UsersRoutes),
    SharedModule,
    FormsModule,
    MatIconModule,
  ]
})
export class UsersModule { }
