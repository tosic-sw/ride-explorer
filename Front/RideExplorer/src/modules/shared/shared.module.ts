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
import { FormsModule } from '@angular/forms';
import { MatSnackBarModule } from '@angular/material/snack-bar';

@NgModule({
  declarations: [
    PaginationComponent,
    SearchInputBtnComponent,
    SelectCustomTextComponent,
    DateFormatPipe
  ],
  imports: [
    CommonModule,
    FormsModule,
    MatSelectModule,
    MatSnackBarModule
  ],
  exports: [
    PaginationComponent,
    SearchInputBtnComponent,
    SelectCustomTextComponent,
    DateFormatPipe
  ],
  providers: [
    SnackBarService,
    UtilService,
    { provide: HTTP_INTERCEPTORS, useClass: Interceptor, multi: true },
  ],
})
export class SharedModule { }
