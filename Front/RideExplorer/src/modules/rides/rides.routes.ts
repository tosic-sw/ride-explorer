import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { CreateRidePageComponent } from "./pages/create-ride-page/create-ride-page.component";
import { DriverFinishedPageComponent } from "./pages/driver-finished-page/driver-finished-page.component";
import { DriverUnfinishedPageComponent } from "./pages/driver-unfinished-page/driver-unfinished-page.component";
import { UpdateRidePageComponent } from "./pages/update-ride-page/update-ride-page.component";
import { ViewRidePageComponent } from "./pages/view-ride-page/view-ride-page.component";

export const RidesRoutes: Routes = [
    {
      path: "create",
      pathMatch: "full",
      component: CreateRidePageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "DRIVER" }
    },
    {
      path: "update/:id",
      pathMatch: "full",
      component: UpdateRidePageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "DRIVER" }
    },
    {
      path: "view/:id",
      pathMatch: "full",
      component: ViewRidePageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN|DRIVER|PASSENGER" }
    },
    {
      path: "finished",
      pathMatch: "full",
      component: DriverFinishedPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "DRIVER" }
    },
    {
      path: "unfinished",
      pathMatch: "full",
      component: DriverUnfinishedPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "DRIVER" }
    },
];