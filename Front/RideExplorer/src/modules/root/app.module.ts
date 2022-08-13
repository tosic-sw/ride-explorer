import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HeaderAdminComponent } from './components/headers/header-admin/header-admin.component';
import { HeaderCommonComponent } from './components/headers/header-common/header-common.component';
import { RootLayoutPageComponent } from './pages/root-layout-page/root-layout-page.component';
import { AuthModule } from '../auth/auth.module';
import { SharedModule } from '../shared/shared.module';
import { HttpClientModule } from '@angular/common/http';
import { ComplaintsModule } from '../complaints/complaints.module';
import { HeaderPassengerComponent } from './components/headers/header-passenger/header-passenger.component';
import { HeaderDriverComponent } from './components/headers/header-driver/header-driver.component';


@NgModule({
  declarations: [
    AppComponent,
    RootLayoutPageComponent,
    HeaderAdminComponent,
    HeaderCommonComponent,
    HeaderPassengerComponent,
    HeaderDriverComponent
  ],
  imports: [
    CommonModule,
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    AuthModule,
    SharedModule,
    HttpClientModule,
    ComplaintsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
