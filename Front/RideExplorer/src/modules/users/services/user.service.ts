import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse, HttpParams } from "@angular/common/http";
import { Observable } from "rxjs";
import { DriverWithCarDTO, UserDTO, UserForUpdateDTO } from "../models/user-dto";
import { MessageResponse } from "src/modules/shared/models/message-response";

@Injectable({
  providedIn: "root",
})
export class UserService {

  private headers = new HttpHeaders({ "Content-Type": "application/json" });
  
  constructor(private http: HttpClient) {}

  search(search: string, page: number, size: number, role: string, verified: boolean = true): Observable<HttpResponse<UserDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: new HttpHeaders({ "Content-Type": 'application/json' }),
      observe: 'response',
      params: new HttpParams()
        .set("search", search)
        .append("page", String(page))
        .append("size", String(size))
    };

    let url: string = "ride-explorer/api/users/search/" + role;

    if(!verified && role === "driver")
      url = url + "/unverified"

    return this.http.get<HttpResponse<UserDTO[]>>(url, queryParams);
  }

  banUser(username: string, role: string): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.put<HttpResponse<MessageResponse>>(`ride-explorer/api/users/ban/${role}/${username}`, queryParams);
  }

  deleteUser(username: string, role: string): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.delete<HttpResponse<MessageResponse>>(`ride-explorer/api/users/${role}/${username}`, queryParams);
  }

  getPassenger(username: string): Observable<HttpResponse<UserDTO>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.get<HttpResponse<UserDTO>>(`ride-explorer/api/users/passenger/${username}`, queryParams);
  }

  getDriver(username: string): Observable<HttpResponse<DriverWithCarDTO>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.get<HttpResponse<DriverWithCarDTO>>(`ride-explorer/api/users/driver/${username}`, queryParams);
  }

  getUserForUpdate(): Observable<HttpResponse<UserForUpdateDTO>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.get<HttpResponse<UserForUpdateDTO>>(`ride-explorer/api/users/profile`, queryParams);
  }

  updateUser(dto: UserForUpdateDTO): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};
    
    queryParams = { 
      headers: this.headers, 
      observe: "response" 
    };

    return this.http.put<HttpResponse<MessageResponse>>(`ride-explorer/api/users/profile`, dto, queryParams);
  }

}
