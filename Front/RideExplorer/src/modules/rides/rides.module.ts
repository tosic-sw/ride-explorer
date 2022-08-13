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

@NgModule({
  declarations: [
    CreateRidePageComponent,
    RideManipulationComponent,
    ViewRidePageComponent,
    UpdateRidePageComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(RidesRoutes),
    SharedModule,
    FormsModule,
    MatIconModule,
  ]
})
export class RidesModule { }
