import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { CreateRidePageComponent } from "./pages/create-ride-page/create-ride-page.component";

export const RidesRoutes: Routes = [
    {
      path: "create",
      pathMatch: "full",
      component: CreateRidePageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "DRIVER" }
    },
];