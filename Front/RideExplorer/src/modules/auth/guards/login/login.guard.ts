import { Injectable } from "@angular/core";
import { Router, CanActivate } from "@angular/router";
import { UtilService } from "src/modules/shared/services/util.service";
import { AuthService } from "../../services/auth-service/auth.service";

@Injectable({
  providedIn: "root",
})
export class LoginGuard implements CanActivate {
  constructor(public auth: AuthService, public router: Router, public utilsService: UtilService) { }

  canActivate(): boolean {
    if (this.auth.isLoggedIn()) {
      let role = this.utilsService.getLoggedUserRole();
      const username = this.utilsService.getLoggedUserUsername();
      if (role === "ADMIN") {
        this.router.navigate(["ridexplorer/users/passengers"]);
      }
      else if (role === "DRIVER") {
        this.router.navigate([`ridexplorer/users/driver/${username}`]);
      }
      else if (role === "PASSENGER") {
        this.router.navigate([`ridexplorer/users/passenger/${username}`]);
      }

      return false;
    }
    return true;
  }
}
