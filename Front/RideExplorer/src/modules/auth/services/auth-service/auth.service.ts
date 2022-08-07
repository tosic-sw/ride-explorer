import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Observable } from "rxjs";
import { Login } from "src/modules/shared/models/login";
import { Token } from "src/modules/shared/models/token";
import { MessageResponse } from "src/modules/shared/models/message-response";
import { DriverRegistrationDTO, RegistrationDTO } from "../../models/registration-dto";

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

  adminPassRegistration(dto: RegistrationDTO, role: string): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.post<HttpResponse<MessageResponse>>("ride-explorer/api/users/registration/" + role, dto, queryParams);
  }

  driverRegistration(dto: DriverRegistrationDTO): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.post<HttpResponse<MessageResponse>>("ride-explorer/api/users/registration/driver", dto, queryParams);
  }

  driverVerification(username: string): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.put<HttpResponse<MessageResponse>>("ride-explorer/api/users/registration/driver/verify/" + username, queryParams);
  }


}
