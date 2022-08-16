import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { UnverifiedReservationsPageComponent } from "./pages/unverified-reservations-page/unverified-reservations-page.component";
import { VerifiedReservationsPageComponent } from "./pages/verified-reservations-page/verified-reservations-page.component";

export const ReservationsRoutes: Routes = [
    {
      path: "verified",
      pathMatch: "full",
      component: VerifiedReservationsPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "PASSENGER" }
    },
    {
        path: "unverified",
        pathMatch: "full",
        component: UnverifiedReservationsPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "PASSENGER" }
      },
];