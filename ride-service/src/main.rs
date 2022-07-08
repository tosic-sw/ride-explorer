#![feature(proc_macro_hygiene, decl_macro)]

extern crate rocket;
extern crate ride_service;
extern crate diesel;

use self::ride_service::*;
use diesel::result::Error;
use ride_service::models::*;
use rocket::{*, http::Status};
use rocket_contrib::json::Json;
use rocket::response::status;

#[get("/drives/<id>")]
pub fn find(id: i32) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();
    find_one(&conn, id)
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[post("/drives", format ="application/json", data = "<new_drive>")]
pub fn create(new_drive: Json<NewDrive>) -> Result<status::Created<Json<Drive>>, Status> {
    let conn = establish_connection();

    create_drive(&conn, &new_drive.into_inner())
        .map(|drive| drive_created(drive))
        .map_err(|error| error_status(error))
}

#[put("/drives", format ="application/json", data = "<dto>")]
pub fn update(dto: Json<UpdateDriveDTO>) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    update_drive(&conn, &dto.into_inner())
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[put("/drives/finish/<id>")]
pub fn finish(id: i32) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    finish_drive(&conn, id)
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[put("/drives/reserve", format ="application/json", data = "<dto>")]
pub fn reserve(dto: Json<ReserveDTO>) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    reserve_drive(&conn, dto.id, dto.places)
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[delete("/drives/<id>")]
pub fn delete(id: i32) -> Result<status::NoContent, Status> {
    let conn = establish_connection();

    delete_drive(&conn, id)
        .map(|_| status::NoContent)
        .map_err(|error| error_status(error))
}

#[post("/drives/search", format ="application/json", data = "<dto>")]
pub fn search(dto: Json<SearchDTO>) -> Result<Json<Vec<Drive>>, Status> {
    let conn = establish_connection();

    search_drives(&conn, &dto.departure_location, &dto.destination)
        .map(|drives| Json(drives))
        .map_err(|error| error_status(error))
}

#[get("/drives")]
pub fn all() -> Result<Json<Vec<Drive>>, Status> {

    let conn = establish_connection();
    get_all(&conn)
        .map(|drives| Json(drives))
        .map_err(|error| error_status(error))
}


fn error_status(error: Error) -> Status {
    match error {
        Error::NotFound => Status::NotFound,
        _ => Status::InternalServerError
    }
}

fn drive_created(drive: Drive) -> status::Created<Json<Drive>> {
    status::Created("".to_string(), Some(Json(drive)))
}

fn main() {
    rocket::ignite().mount("/api", routes![find, create, update, finish, reserve, delete, search, all],).launch();
}
