import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse, HttpParams } from "@angular/common/http";
import { Observable } from "rxjs";
import { RatingDTO, ViewRatingDTO } from "../models/rating-dto";
import { MessageResponse } from "../models/message-response";

@Injectable({
  providedIn: "root",
})
export class RatingService {

  private headers = new HttpHeaders({ "Content-Type": "application/json" });
  
  constructor(private http: HttpClient) {}

  getRatings(username:string, page: number, size: number): Observable<HttpResponse<ViewRatingDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response',
      params: new HttpParams()
        .set("page", String(page))
        .append("size", String(size))
    };

    let url: string = `ride-explorer/api/ratings/evaluated/${username}`;

    return this.http.get<HttpResponse<ViewRatingDTO[]>>(url, queryParams);
  }

  createRating(dto: RatingDTO): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/ratings`;

    return this.http.post<HttpResponse<MessageResponse>>(url, dto, queryParams);
  }

  updateRating(dto: RatingDTO): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/ratings`;

    return this.http.put<HttpResponse<MessageResponse>>(url, dto, queryParams);
  }

  deleteRating(id: number): Observable<HttpResponse<MessageResponse>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response'
    };

    let url: string = `ride-explorer/api/ratings/${id}`;

    return this.http.post<HttpResponse<MessageResponse>>(url, queryParams);
  }

}
