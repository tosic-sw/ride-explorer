use serde:: { Serialize, Deserialize };
use super::schema::drives;

#[derive(Queryable, Serialize)]
pub struct Drive {
    pub id: i32,
    pub departure_location: String,
    pub destination: String,
    pub departure_date_time: i64,
    pub departure_address: String,
    pub free_places: i32,
    pub planned_arrival_time: i64,
    pub note: String,
    pub finished: bool,
    pub distance: i32,
}

#[derive(Insertable, Deserialize)]
#[table_name="drives"]
pub struct NewDrive {
    pub departure_location: String,
    pub destination: String,
    pub departure_date_time: i64,
    pub departure_address: String,
    pub free_places: i32,
    pub planned_arrival_time: i64,
    pub note: String,
    pub distance: i32,
}

#[derive(Deserialize)]
pub struct UpdateDriveDTO {
    pub id: i32,
    pub departure_address: String,
    pub free_places: i32,
    pub note: String,
}

#[derive(Deserialize)]
pub struct ReserveDTO {
    pub id: i32,
    pub places: i32,
}

#[derive(Serialize)]
pub struct StatusMessage {
    pub message: String,
}

#[derive(Serialize)]
pub struct Drives {
    pub drives: Vec<Drive>
}


#[derive(Deserialize)]
pub struct SearchDTO {
    pub departure_location: String,
    pub destination: String,
}
