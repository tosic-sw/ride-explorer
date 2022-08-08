import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { AdminsPageComponent } from "./pages/admins-page/admins-page.component";
import { DriversPageComponent } from "./pages/drivers-page/drivers-page.component";
import { PassengersPageComponent } from "./pages/passengers-page/passengers-page.component";
import { UnverifiedDriversPageComponent } from "./pages/unverified-drivers-page/unverified-drivers-page.component";


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
      // canActivate: [RoleGuard],
      // data: { expectedRoles: "ADMIN" }
    },
    {
      path: "admins",
      pathMatch: "full",
      component: AdminsPageComponent,
      // canActivate: [RoleGuard],
      // data: { expectedRoles: "ADMIN" }
    },
    {
      path: "passengers",
      pathMatch: "full",
      component: PassengersPageComponent,
      // canActivate: [RoleGuard],
      // data: { expectedRoles: "ADMIN" }
    },
  ];