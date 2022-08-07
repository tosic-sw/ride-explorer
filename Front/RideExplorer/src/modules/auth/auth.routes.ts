import { Routes } from "@angular/router";
import { LoginGuard } from "./guards/login/login.guard";
import { RoleGuard } from "./guards/role/role.guard";
import { AdminRegistrationPageComponent } from "./pages/admin-registration-page/admin-registration-page.component";
import { DriverRegistrationPageComponent } from "./pages/driver-registration-page/driver-registration-page.component";
import { DriverVerificationPageComponent } from "./pages/driver-verification-page/driver-verification-page.component";
import { LoginComponent } from "./pages/login/login.component";
import { PassengerRegistrationPageComponent } from "./pages/passenger-registration-page/passenger-registration-page.component";

export const AuthRoutes: Routes = [
  {
    path: "login",
    pathMatch: "full",
    component: LoginComponent,
    canActivate: [LoginGuard],
  },
  {
    path: "passenger-registration",
    pathMatch: "full",
    component: PassengerRegistrationPageComponent,
    canActivate: [LoginGuard],
  },
  {
    path: "admin-registration",
    pathMatch: "full",
    component: AdminRegistrationPageComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "ADMIN" }
  },
  {
    path: "driver-registration",
    pathMatch: "full",
    component: DriverRegistrationPageComponent,
    canActivate: [LoginGuard],
  },
  {
    path: "driver-verification/:username",
    pathMatch: "full",
    component: DriverVerificationPageComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "ADMIN" }
  },
];
