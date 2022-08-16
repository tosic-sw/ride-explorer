import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse, HttpParams } from "@angular/common/http";
import { Observable } from "rxjs";
import { CreateReservationDTO, ReservationDTO } from "src/modules/shared/models/reservation-dtos";
import { MessageResponse } from "src/modules/shared/models/message-response";

@Injectable({
  providedIn: "root",
})
export class ReservationService {

  private headers = new HttpHeaders({ "Content-Type": "application/json" });
  
  constructor(private http: HttpClient) {}

  createReservation(dto: CreateReservationDTO): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/reservations`;

    return this.http.post<HttpResponse<MessageResponse>>(url, dto, queryParams);
  }

  getVerifiedForDrive(driveId: number, page: number, size: number): Observable<HttpResponse<ReservationDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response',
      params: new HttpParams()
        .set("page", String(page))
        .append("size", String(size))
    };

    let url: string = `ride-explorer/api/reservations/drive/${driveId}/verified`;

    return this.http.get<HttpResponse<ReservationDTO[]>>(url, queryParams);
  }

  getVerifiedForDriveAndDriver(driveId: number, page: number, size: number): Observable<HttpResponse<ReservationDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response',
      params: new HttpParams()
        .set("page", String(page))
        .append("size", String(size))
    };

    let url: string = `ride-explorer/api/reservations/driver/${driveId}/verified`;

    return this.http.get<HttpResponse<ReservationDTO[]>>(url, queryParams);
  }

  getUnverifiedForDriveAndDriver(driveId: number, page: number, size: number): Observable<HttpResponse<ReservationDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response',
      params: new HttpParams()
        .set("page", String(page))
        .append("size", String(size))
    };

    let url: string = `ride-explorer/api/reservations/driver/${driveId}/unverified`;

    return this.http.get<HttpResponse<ReservationDTO[]>>(url, queryParams);
  }

  getVerifiedForPassenger(page: number, size: number): Observable<HttpResponse<ReservationDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response',
      params: new HttpParams()
        .set("page", String(page))
        .append("size", String(size))
    };

    let url: string = `ride-explorer/api/reservations/user/verified`;

    return this.http.get<HttpResponse<ReservationDTO[]>>(url, queryParams);
  }

  getUnverifiedForPassenger(page: number, size: number): Observable<HttpResponse<ReservationDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response',
      params: new HttpParams()
        .set("page", String(page))
        .append("size", String(size))
    };

    let url: string = `ride-explorer/api/reservations/user/unverified`;

    return this.http.get<HttpResponse<ReservationDTO[]>>(url, queryParams);
  }

}