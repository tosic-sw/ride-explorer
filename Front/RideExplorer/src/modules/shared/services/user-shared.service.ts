import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Observable } from "rxjs";
import { DriverWithCarDTO } from "../models/driver-shared-dto";

@Injectable({
  providedIn: "root",
})
export class UserSharedService {

  private headers = new HttpHeaders({ "Content-Type": "application/json" });
  
  constructor(private http: HttpClient) {}

  getUnverifiedDriver(username: string): Observable<HttpResponse<DriverWithCarDTO>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.get<HttpResponse<DriverWithCarDTO>>("ride-explorer/api/users/driver/unverified/" + username, queryParams);
  }


}
