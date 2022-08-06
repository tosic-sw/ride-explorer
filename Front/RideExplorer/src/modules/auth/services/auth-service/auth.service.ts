import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { Login } from "src/modules/shared/models/login";
import { Token } from "src/modules/shared/models/token";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private headers = new HttpHeaders({ "Content-Type": "application/json" });

  constructor(private http: HttpClient) {}

  login(auth: Login): Observable<Token> {
    return this.http.post<Token>("ride-explorer/api/users/login", auth, {
      headers: this.headers,
      responseType: "json",
    });
  }

  logout(): void {
    sessionStorage.removeItem("user");
  }

  isLoggedIn(): boolean {
    if (!sessionStorage.getItem("user")) {
      return false;
    }
    return true;
  }
}
