import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable({
  providedIn: 'root'
})
export class UtilService {

  constructor(private http: HttpClient, private router: Router) { }

  public getNoPages(totalItems: number, pageSize: number): number {
    return Math.ceil(totalItems / pageSize);
  }

  public getLoggedUserRole(): string {
    const item = sessionStorage.getItem("user");

    if (item) {
      const jwt: JwtHelperService = new JwtHelperService();
      return jwt.decodeToken(item).role;
    }
    return "";
  }

  public getLoggedUserRoleLower(): string {
    const item = sessionStorage.getItem("user");

    if (item) {
      const jwt: JwtHelperService = new JwtHelperService();
      return jwt.decodeToken(item).role.toLowerCase();
    }
    return "";
  }

  public getLoggedUserUsername(): string {
    const item = sessionStorage.getItem("user");

    if (item) {
      const jwt: JwtHelperService = new JwtHelperService();
      return jwt.decodeToken(item).username;
    }
    return "";
  }

  public navigateToMyProfile() {
    const role = this.getLoggedUserRole();
    const username = this.getLoggedUserUsername();

    if (role === "ADMIN") {
      this.router.navigate(["ridexplorer/users/passengers"]);
    }
    else if (role === "DRIVER") {
      this.router.navigate([`ridexplorer/users/driver/${username}`]);
    }
    else if (role === "PASSENGER") {
      this.router.navigate([`ridexplorer/users/passenger/${username}`]);
    }
  }

  public isNumber(str: string): boolean {
    if (typeof str !== 'string') {
      return false;
    }
  
    if (str.trim() === '') {
      return false;
    }
  
    return !Number.isNaN(Number(str));
  }

}
