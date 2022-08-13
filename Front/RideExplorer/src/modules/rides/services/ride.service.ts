import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Observable } from "rxjs";
import { DriveDTO, NewDriveDTO, ReserveDTO, SearchDTO, UpdateDriveDTO } from "../models/drive-dto";

@Injectable({
  providedIn: "root",
})
export class RideService {

  private headers = new HttpHeaders({ "Content-Type": "application/json" });
  
  constructor(private http: HttpClient) {}

  searchRides(searchDTO: SearchDTO): Observable<HttpResponse<DriveDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives`;

    return this.http.post<HttpResponse<DriveDTO[]>>(url, searchDTO, queryParams);
  }

  createRide(dto: NewDriveDTO): Observable<HttpResponse<DriveDTO>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives`;

    return this.http.post<HttpResponse<DriveDTO>>(url, dto, queryParams);
  }

  getRide(id: number): Observable<HttpResponse<DriveDTO>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/${id}`;

    return this.http.get<HttpResponse<DriveDTO>>(url, queryParams);
  }

  updateRide(dto: UpdateDriveDTO): Observable<HttpResponse<DriveDTO>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives`;

    return this.http.put<HttpResponse<DriveDTO>>(url, dto, queryParams);
  }

  finishRide(username: string, id: number): Observable<HttpResponse<DriveDTO>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/driver/${username}/finish/${id}`;

    return this.http.put<HttpResponse<DriveDTO>>(url, queryParams);
  }

  reservePlace(dto: ReserveDTO): Observable<HttpResponse<DriveDTO>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/adjust-places`;

    return this.http.put<HttpResponse<DriveDTO>>(url, dto, queryParams);
  }

  deleteDrive(username: string, id: number): Observable<HttpResponse<any>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/driver/${username}/${id}`;

    return this.http.delete<HttpResponse<any[]>>(url, queryParams);
  }


}