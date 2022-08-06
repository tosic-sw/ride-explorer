import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthRoutes } from './auth.routes';
import { LoginComponent } from './pages/login/login.component';
import { ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';



@NgModule({
  declarations: [LoginComponent],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(AuthRoutes),
  ],
  providers: []
})
export class AuthModule { }
