#[macro_use]
extern crate diesel;
extern crate dotenv;

pub mod schema;
pub mod models;

use diesel::{prelude::*};
use diesel::pg::PgConnection;

use dotenv::dotenv;
use std::env;
use std::time::{SystemTime, UNIX_EPOCH};

use crate::schema::drives::{departure_address, free_places, note, finished};
use self::models::{Drive, NewDrive, UpdateDriveDTO};

pub fn establish_connection() -> PgConnection {
    dotenv().ok();

    let database_url = env::var("DATABASE_URL")
        .expect("DATABASE_URL must be set");
    PgConnection::establish(&database_url)
        .expect(&format!("Error connecting to {}", database_url))
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

pub fn finish_drive(conn: &PgConnection, id: i32) -> QueryResult<Drive> {
    use schema::drives::dsl::drives;

    let drive = diesel::update(drives.find(id))
        .set(finished.eq(true))
        .get_result::<Drive>(conn);

    drive
}

pub fn reserve_drive(conn: &PgConnection, id: i32, reserved_num: i32) -> QueryResult<Drive> {
    use schema::drives::dsl::drives;

    let drive = diesel::update(drives.find(id))
        .set(free_places.eq(free_places - reserved_num))
        .get_result::<Drive>(conn);

    drive
}

pub fn delete_drive(conn: &PgConnection, _id: i32) -> QueryResult<usize> {
    use schema::drives::dsl::*;
    diesel::delete(drives.filter(id.eq(_id))).execute(conn)
}

pub fn find_one(conn: &PgConnection, _id: i32) -> QueryResult<Drive> {
    use schema::drives;

    drives::table.find(_id).get_result::<Drive>(conn)
}

pub fn search_drives(conn: &PgConnection, depart: &String, dest: &String) -> QueryResult<Vec<Drive>> {
    use schema::drives::dsl::*;

    let start = SystemTime::now();
    let since_the_epoch = start
        .duration_since(UNIX_EPOCH)
        .expect("Time went backwards");
    let millis: i64 = since_the_epoch.as_millis().try_into().unwrap();

    let results = drives
        .filter(departure_location.eq(&depart))
        .filter(destination.eq(dest))
        .filter(finished.eq(false))
        .filter(departure_date_time.gt(millis))
        .limit(100)
        .load::<Drive>(conn);

    results
}

pub fn get_all(conn: &PgConnection) -> QueryResult<Vec<Drive>> {
    use schema::drives::dsl::*;

    let results = drives
        .limit(500)
        .load::<Drive>(conn);

    results
}