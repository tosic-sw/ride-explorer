import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateRidePageComponent } from './pages/create-ride-page/create-ride-page.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SharedModule } from '../shared/shared.module';
import { MatIconModule } from '@angular/material/icon';
import { RouterModule } from '@angular/router';
import { RidesRoutes } from './rides.routes';
import { RideManipulationComponent } from './components/ride-manipulation/ride-manipulation.component';
import { ViewRidePageComponent } from './pages/view-ride-page/view-ride-page.component';
import { UpdateRidePageComponent } from './pages/update-ride-page/update-ride-page.component';
import { DriverFinishedPageComponent } from './pages/driver-finished-page/driver-finished-page.component';
import { DriverUnfinishedPageComponent } from './pages/driver-unfinished-page/driver-unfinished-page.component';
import { RideTableComponent } from './components/ride-table/ride-table.component';

@NgModule({
  declarations: [
    CreateRidePageComponent,
    RideManipulationComponent,
    ViewRidePageComponent,
    UpdateRidePageComponent,
    DriverFinishedPageComponent,
    DriverUnfinishedPageComponent,
    RideTableComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(RidesRoutes),
    SharedModule,
    FormsModule,
    MatIconModule
  ]
})
export class RidesModule { }
