import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PaginationComponent } from './components/pagination/pagination.component';
import { SearchInputBtnComponent } from './components/search-input-btn/search-input-btn.component';
import { SelectCustomTextComponent } from './components/select-custom-text/select-custom-text.component';
import { DateFormatPipe } from './pipes/date-format.pipe';
import { SnackBarService } from './services/snack-bar.service';
import { UtilService } from './services/util.service';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { Interceptor } from './interceptors/interceptor.interceptor';
import {MatSelectModule} from '@angular/material/select';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { CreateUpdateProfileComponent } from './components/create-update-profile/create-update-profile.component';
import { ReservationTableComponent } from './components/reservation-table/reservation-table.component';
import { ReservationService } from './services/reservation.service';
import { ReservationsPassengerTableComponent } from './components/reservations-passenger-table/reservations-passenger-table.component';
import { MatIconModule } from '@angular/material/icon';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatTooltipModule } from '@angular/material/tooltip';

@NgModule({
  declarations: [
    PaginationComponent,
    SearchInputBtnComponent,
    SelectCustomTextComponent,
    DateFormatPipe,
    CreateUpdateProfileComponent,
    ReservationTableComponent,
    ReservationsPassengerTableComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    MatSelectModule,
    MatSnackBarModule,
    FormsModule,
    ReactiveFormsModule,
    MatIconModule,
    MatTooltipModule
  ],
  exports: [
    PaginationComponent,
    SearchInputBtnComponent,
    SelectCustomTextComponent,
    DateFormatPipe,
    CreateUpdateProfileComponent,
    ReservationTableComponent,
    ReservationsPassengerTableComponent
  ],
  providers: [
    SnackBarService,
    UtilService,
    ReservationService,
    { provide: HTTP_INTERCEPTORS, useClass: Interceptor, multi: true },
  ],
})
export class SharedModule { }
