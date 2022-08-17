#[macro_use]
extern crate diesel;
extern crate dotenv;

pub mod schema;
pub mod models;

use diesel::{prelude::*};
use diesel::pg::PgConnection;

use dotenv::dotenv;
use models::PageableDTO;
use std::env;
use std::time::{SystemTime, UNIX_EPOCH};

use crate::schema::drives::{departure_address, free_places, note};
use self::models::{Drive, NewDrive, UpdateDriveDTO, SearchDTO};

pub fn establish_connection() -> PgConnection {
    dotenv().ok();

    let database_url = env::var("DATABASE_URL")
        .expect("DATABASE_URL must be set");
    PgConnection::establish(&database_url)
        .expect(&format!("Error connecting to {}", database_url))
}

pub fn insert_d1(conn: &PgConnection) -> QueryResult<Drive> {
    let d1 = &NewDrive {
        driver_username: String::from("tica"),
        departure_location: String::from("novi sad"),
        destination: String::from("beograd"),
        departure_date_time: 1662814800000,
        departure_address: String::from("Misiceva 3"),
        free_places: 3,
        planned_arrival_time: 1662818400000,
        note: String::from("Ciao bella"),
        distance: 90,
    };

    use schema::drives;

    diesel::insert_into(drives::table)
        .values(d1)
        .get_result(conn)
}

pub fn insert_d2(conn: &PgConnection) -> QueryResult<Drive> {
    let d2 = &NewDrive {
        driver_username: String::from("tica"),
        departure_location: String::from("novi sad"),
        destination: String::from("beograd"),
        departure_date_time: 1662814800000,
        departure_address: String::from("Bojoviceva 4"),
        free_places: 3,
        planned_arrival_time: 1662818400000,
        note: String::from("Ciao bella"),
        distance: 90,
    };

    use schema::drives;

    diesel::insert_into(drives::table)
        .values(d2)
        .get_result(conn)
}

pub fn insert_d3(conn: &PgConnection) -> QueryResult<Drive> {
    let d3 = &NewDrive {
        driver_username: String::from("ukica"),
        departure_location: String::from("novi sad"),
        destination: String::from("beograd"),
        departure_date_time: 1662814800000,
        departure_address: String::from("Putnikova 5"),
        free_places: 3,
        planned_arrival_time: 1662818400000,
        note: String::from("Ciao bella"),
        distance: 90,
    };

    use schema::drives;

    diesel::insert_into(drives::table)
        .values(d3)
        .get_result(conn)
}

pub fn create_drive(conn: &PgConnection, new_drive: &NewDrive) -> QueryResult<Drive> {
    use schema::drives;

    diesel::insert_into(drives::table)
        .values(new_drive)
        .get_result(conn)
}

pub fn update_drive(conn: &PgConnection, dto: &UpdateDriveDTO) -> QueryResult<Drive> {
    use schema::drives::dsl::drives;

    let drive = diesel::update(drives.find(dto.id))
        .set((departure_address.eq(&dto.departure_address), free_places.eq(dto.free_places), note.eq(&dto.note)))
        .get_result::<Drive>(conn);

    drive
}

pub fn finish_drive(conn: &PgConnection, username: &String, _id: i32) -> QueryResult<Drive> {
    use schema::drives::dsl::*;

    let drive = diesel::update
        (drives.
            filter(id.eq(_id)).
            filter(finished.eq(false)).
            filter(driver_username.eq(username)))
        .set(finished.eq(true))
        .get_result::<Drive>(conn);

    drive
}

pub fn reserve_drive(conn: &PgConnection, _id: i32, reserved_num: i32) -> QueryResult<Drive> {
    use schema::drives::dsl::*;

    let drive = diesel::update
        (drives.
            filter((free_places-reserved_num).gt(-1)).
            filter(finished.eq(false)).
            filter(id.eq(_id)))
        .set(free_places.eq(free_places - reserved_num))
        .get_result::<Drive>(conn);

    drive
}

pub fn delete_drive(conn: &PgConnection, driver: &String, _id: i32) -> QueryResult<usize> {
    use schema::drives::dsl::*;

    diesel::delete(drives
        .filter(id.eq(_id))
        .filter(finished.eq(true))
        .filter(driver_username.eq(driver)))
        .execute(conn)
}

pub fn find_one(conn: &PgConnection, _id: i32) -> QueryResult<Drive> {
    use schema::drives;

    drives::table.find(_id).get_result::<Drive>(conn)
}

pub fn search_drives(conn: &PgConnection, dto: &SearchDTO) -> QueryResult<Vec<Drive>> {
    use schema::drives::dsl::*;

    let offset: i64 = (dto.page - 1) * dto.size;
    let millis: i64 = get_now_since_epoch();

    let results = drives
        .filter(departure_location.eq(&dto.departure_location.trim().to_lowercase()))
        .filter(destination.eq(&dto.destination.trim().to_lowercase()))
        .filter(finished.eq(false))
        .filter(departure_date_time.gt(millis))
        .limit(dto.size)
        .offset(offset)
        .load::<Drive>(conn);

    results
}

pub fn count_search_drives(conn: &PgConnection, dto: &SearchDTO) -> i64 {
    use schema::drives::dsl::*;

    let millis: i64 = get_now_since_epoch();

    let total_elements = drives
        .filter(departure_location.eq(&dto.departure_location.trim().to_lowercase()))
        .filter(destination.eq(&dto.destination.trim().to_lowercase()))
        .filter(finished.eq(false))
        .filter(departure_date_time.gt(millis))
        .count()
        .get_result(conn);

    total_elements.unwrap()
}

pub fn find_one_finished_driver(conn: &PgConnection, driver: &String, _id: i32) -> QueryResult<Drive> {
    use schema::drives::dsl::*;

    let results = drives
        .filter(finished.eq(true))
        .filter(id.eq(_id))
        .filter(driver_username.eq(driver))
        .get_result::<Drive>(conn);

    results
}

pub fn find_one_unfinished_driver(conn: &PgConnection, driver: &String, _id: i32) -> QueryResult<Drive> {
    use schema::drives::dsl::*;

    let results = drives
        .filter(finished.eq(false))
        .filter(id.eq(_id))
        .filter(driver_username.eq(driver))
        .get_result::<Drive>(conn);

    results
}

pub fn find_one_finished(conn: &PgConnection, _id: i32) -> QueryResult<Drive> {
    use schema::drives::dsl::*;

    let results = drives
        .filter(finished.eq(true))
        .filter(id.eq(_id))
        .get_result::<Drive>(conn);

    results
}

pub fn find_finished_driver(conn: &PgConnection, driver: &String, dto: &PageableDTO) -> QueryResult<Vec<Drive>> {
    use schema::drives::dsl::*;

    let offset: i64 = (dto.page - 1) * dto.size;

    let results = drives
        .filter(finished.eq(true))
        .filter(driver_username.eq(driver))
        .offset(offset)
        .limit(dto.size)
        .load::<Drive>(conn);

    results
}

pub fn count_finished_driver(conn: &PgConnection, driver: &String) -> i64 {
    use schema::drives::dsl::*;

    let count = drives
        .filter(finished.eq(true))
        .filter(driver_username.eq(driver))
        .count()
        .get_result(conn);

    count.unwrap()
}

pub fn find_unfinished_driver(conn: &PgConnection, driver: &String, dto: &PageableDTO) -> QueryResult<Vec<Drive>> {
    use schema::drives::dsl::*;

    let offset: i64 = (dto.page - 1) * dto.size;

    let results = drives
        .filter(finished.eq(false))
        .filter(driver_username.eq(driver))
        .offset(offset)
        .limit(dto.size)
        .load::<Drive>(conn);

    results
}

pub fn count_unfinished_driver(conn: &PgConnection, driver: &String) -> i64 {
    use schema::drives::dsl::*;

    let count = drives
        .filter(finished.eq(false))
        .filter(driver_username.eq(driver))
        .count()
        .get_result(conn);

    count.unwrap()
}

pub fn get_all(conn: &PgConnection) -> QueryResult<Vec<Drive>> {
    use schema::drives::dsl::*;

    let results = drives
        .limit(500)
        .load::<Drive>(conn);

    results
}

fn get_now_since_epoch() -> i64 {
    let start = SystemTime::now();
    let since_the_epoch = start
        .duration_since(UNIX_EPOCH)
        .expect("Time went backwards");
    let millis: i64 = since_the_epoch.as_millis().try_into().unwrap();

    return millis;
}