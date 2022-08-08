import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NotFoundPageComponent } from './pages/not-found-page/not-found-page.component';
import { RootLayoutPageComponent } from './pages/root-layout-page/root-layout-page.component';

const routes: Routes = [
  {
    path: "ridexplorer",
    component: RootLayoutPageComponent,
    children: [
      {
        path: "auth",
        loadChildren: () =>
          import("./../auth/auth.module").then((m) => m.AuthModule),
      },
      {
        path: "users",
        loadChildren: () =>
          import("./../users/users.module").then((m) => m.UsersModule),
      },
    ],
  },
  {
    path: "",
    redirectTo: "ridexplorer/auth/login",
    pathMatch: "full",
  },
  { 
    path: "**", 
    component: NotFoundPageComponent 
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
