import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse, HttpParams } from "@angular/common/http";
import { Observable } from "rxjs";
import { ComplaintDTO, CreateComplaintDTO } from "../../shared/models/complaint-dto";
import { MessageResponse } from "../models/message-response";

@Injectable({
  providedIn: "root",
})
export class ComplaintService {

  private headers = new HttpHeaders({ "Content-Type": "application/json" });
  
  constructor(private http: HttpClient) {}

  getComplaints(page: number, size: number): Observable<HttpResponse<ComplaintDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response',
      params: new HttpParams()
        .set("page", String(page))
        .append("size", String(size))
    };

    let url: string = `ride-explorer/api/complaints`;

    return this.http.get<HttpResponse<ComplaintDTO[]>>(url, queryParams);
  }

  createComplaint(dto: CreateComplaintDTO): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/complaints`;

    return this.http.post<HttpResponse<MessageResponse>>(url, dto, queryParams);
  }

}