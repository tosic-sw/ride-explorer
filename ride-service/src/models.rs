use serde:: { Serialize, Deserialize };
use super::schema::drives;

#[derive(Queryable, Serialize)]
pub struct Drive {
    pub id: i32,
    pub driver_username: String,
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
    pub driver_username: String,
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
    pub drives: Vec<Drive>,
    pub total_elements: i64,
}


#[derive(Deserialize)]
pub struct SearchDTO {
    pub departure_location: String,
    pub destination: String,    
    pub page: i64,
    pub size: i64,
}

#[derive(Deserialize)]
pub struct PageableDTO {  
    pub page: i64,
    pub size: i64,
}


pub struct Node {
    pub departure_location: String,
    pub destination: String,
    pub drive_ids: Vec<i32>,
    pub distances: Vec<u64>,
    pub departure_date_times: Vec<u64>,
    pub planned_arival_times: Vec<u64>,
}

