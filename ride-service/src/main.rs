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

#[get("/drives/unfinished/<id>/<username>")]
pub fn one_unfinished_driver(id: i32, username: String) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    find_one_unfinished_driver(&conn, &username, id)
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[get("/drives/finished/<id>/<username>")]
pub fn one_finished_driver(id: i32, username: String) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    find_one_finished_driver(&conn, &username, id)
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[get("/drives/finished/<id>")]
pub fn one_finished(id: i32) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    find_one_finished(&conn, id)
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[post("/drives", data = "<new_drive>")]
pub fn create(new_drive: Json<NewDrive>) -> Result<status::Created<Json<Drive>>, Status> {
    let conn = establish_connection();

    create_drive(&conn, &new_drive.into_inner())
        .map(|drive| drive_created(drive))
        .map_err(|error| error_status(error))
}

#[put("/drives", data = "<dto>")]
pub fn update(dto: Json<UpdateDriveDTO>) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    update_drive(&conn, &dto.into_inner())
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[put("/drives/driver/<username>/finish/<id>")]
pub fn finish(username: String, id: i32) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    finish_drive(&conn, &username, id)
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[put("/drives/adjust-places", data = "<dto>")]
pub fn reserve(dto: Json<ReserveDTO>) -> Result<Json<Drive>, Status> {
    let conn = establish_connection();

    reserve_drive(&conn, dto.id, dto.places)
        .map(|drive| Json(drive))
        .map_err(|error| error_status(error))
}

#[delete("/drives/driver/<username>/<id>")]
pub fn delete(username: String, id: i32) -> Result<status::NoContent, Status> {
    let conn = establish_connection();

    let mut can_delete = false;

    let _ = find_one_finished_driver(&conn, &username, id)
        .map(|_| can_delete = true)
        .map_err(|error| error_status(error));

    if can_delete {
        delete_drive(&conn, &username, id)
            .map(|_| status::NoContent)
            .map_err(|error| error_status(error))
    } else {
        Err(error_status(Error::NotFound))
    }

}

#[post("/drives/search", data = "<dto>")]
pub fn search(dto: Json<SearchDTO>) -> Result<Json<Drives>, Status> {
    let conn = establish_connection();
    
    let total_elements = count_search_drives(&conn, &dto);

    search_drives(&conn, &dto)
        .map(|drives| Json(
            Drives {
                drives: drives,
                total_elements: total_elements
            }
        ))
        .map_err(|error| error_status(error))
}

#[post("/drives/driver/finished/<username>", data = "<dto>")]
pub fn finished_driver(username: String, dto: Json<PageableDTO>) -> Result<Json<Drives>, Status> {
    let conn = establish_connection();

    let total_elements = count_finished_driver(&conn, &username);

    find_finished_driver(&conn, &username, &dto)
        .map(|drives| Json(
            Drives {
                drives: drives,
                total_elements: total_elements
            }
        ))
        .map_err(|error| error_status(error))
}

#[post("/drives/driver/unfinished/<username>", data = "<dto>")]
pub fn unfinished_driver(username: String, dto: Json<PageableDTO>) -> Result<Json<Drives>, Status> {
    let conn = establish_connection();

    let total_elements = count_unfinished_driver(&conn, &username);

    find_unfinished_driver(&conn, &username, &dto)
        .map(|drives| Json(
            Drives {
                drives: drives,
                total_elements: total_elements
            }
        ))
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
    let conn = establish_connection();
    _ = insert_d1(&conn);
    _ = insert_d2(&conn);
    _ = insert_d3(&conn);

    rocket::ignite().mount("/api", routes![find, create, update, finish, reserve, delete, search, all, finished_driver, unfinished_driver, one_unfinished_driver,
                                                 one_finished, one_finished_driver],).launch();
}
