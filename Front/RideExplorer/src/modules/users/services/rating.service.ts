import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse, HttpParams } from "@angular/common/http";
import { Observable } from "rxjs";
import { ViewRatingDTO } from "../models/rating-dto";

@Injectable({
  providedIn: "root",
})
export class RatingService {

  private headers = new HttpHeaders({ "Content-Type": "application/json" });
  
  constructor(private http: HttpClient) {}

  getRatings(username:string, page: number, size: number): Observable<HttpResponse<ViewRatingDTO[]>> {
    let queryParams = {};

    queryParams = {
      headers: new HttpHeaders({ "Content-Type": 'application/json' }),
      observe: 'response',
      params: new HttpParams()
        .set("page", String(page))
        .append("size", String(size))
    };

    let url: string = `ride-explorer/api/ratings/evaluated/${username}`;

    return this.http.get<HttpResponse<ViewRatingDTO[]>>(url, queryParams);
  }

}
