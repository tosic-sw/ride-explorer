import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { AdminsPageComponent } from "./pages/admins-page/admins-page.component";
import { DriverPageComponent } from "./pages/driver-page/driver-page.component";
import { DriversPageComponent } from "./pages/drivers-page/drivers-page.component";
import { PassengerPageComponent } from "./pages/passenger-page/passenger-page.component";
import { PassengersPageComponent } from "./pages/passengers-page/passengers-page.component";
import { UnverifiedDriversPageComponent } from "./pages/unverified-drivers-page/unverified-drivers-page.component";
import { UpdateProfilePageComponent } from "./pages/update-profile-page/update-profile-page.component";


export const UsersRoutes: Routes = [
    {
      path: "unverified-drivers",
      pathMatch: "full",
      component: UnverifiedDriversPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN" }
    },
    {
      path: "drivers",
      pathMatch: "full",
      component: DriversPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN" }
    },
    {
      path: "admins",
      pathMatch: "full",
      component: AdminsPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN" }
    },
    {
      path: "passengers",
      pathMatch: "full",
      component: PassengersPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN" }
    },
    {
      path: "driver/:username",
      pathMatch: "full",
      component: DriverPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN|PASSENGER|DRIVER" }
    },
    {
      path: "passenger/:username",
      pathMatch: "full",
      component: PassengerPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN|PASSENGER|DRIVER" }
    },
    {
      path: "update-profile",
      pathMatch: "full",
      component: UpdateProfilePageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN|PASSENGER|DRIVER" }
    },
];