import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { ComplaintsPageComponent } from "./pages/complaints-page/complaints-page.component";

export const ComplaintsRoutes: Routes = [
    {
      path: "all",
      pathMatch: "full",
      component: ComplaintsPageComponent,
      canActivate: [RoleGuard],
      data: { expectedRoles: "ADMIN" }
    },
];