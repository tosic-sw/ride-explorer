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
import { RideSearchTableComponent } from './components/ride-search-table/ride-search-table.component';
import { SearchRidePageComponent } from './pages/search-ride-page/search-ride-page.component';
import { SearchFormComponent } from './components/search-form/search-form.component';
import { MatTooltipModule } from '@angular/material/tooltip';

@NgModule({
  declarations: [
    CreateRidePageComponent,
    RideManipulationComponent,
    ViewRidePageComponent,
    UpdateRidePageComponent,
    DriverFinishedPageComponent,
    DriverUnfinishedPageComponent,
    RideTableComponent,
    RideSearchTableComponent,
    SearchRidePageComponent,
    SearchFormComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(RidesRoutes),
    SharedModule,
    FormsModule,
    MatIconModule,
    MatTooltipModule
  ]
})
export class RidesModule { }
