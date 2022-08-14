import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Observable } from "rxjs";
import { DriveDTO, Drives, NewDriveDTO, PageableDTO, ReserveDTO, SearchDTO, UpdateDriveDTO } from "../models/drive-dto";

@Injectable({
  providedIn: "root",
})
export class RideService {

  private headers = new HttpHeaders({ "Content-Type": "application/json" });
  
  constructor(private http: HttpClient) {}

  searchRides(departureLocation: string, destination: string, page: number, size: number): Observable<HttpResponse<Drives>> { // Izmeni !!!
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/search`;
    const searchDTO: SearchDTO = {
      departure_location: departureLocation,
      destination: destination,
      page: page + 1,
      size: size
    };

    return this.http.post<HttpResponse<Drives>>(url, searchDTO, queryParams);
  }

  getFinishedRidesDriver(username: string, page: number, size: number): Observable<HttpResponse<Drives>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/driver/finished/${username}`;
    const pageableDTO: PageableDTO = {
      page: page + 1,
      size: size
    };
    console.log(pageableDTO);

    return this.http.post<HttpResponse<Drives>>(url, pageableDTO, queryParams);
  }

  getUnfinishedRidesDriver(username: string, page: number, size: number): Observable<HttpResponse<Drives>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/driver/unfinished/${username}`;
    const pageableDTO: PageableDTO = {
      page: page + 1,
      size: size
    };

    return this.http.post<HttpResponse<Drives>>(url, pageableDTO, queryParams);
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

  getRideUnfinishedDriver(id: number, username: string): Observable<HttpResponse<DriveDTO>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/unfinished/${id}/${username}`;

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

  deleteDrive(username: string, id: number): Observable<HttpResponse<any>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/drives/driver/${username}/${id}`;

    return this.http.delete<HttpResponse<any>>(url, queryParams);
  }


}