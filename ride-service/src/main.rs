#![feature(proc_macro_hygiene, decl_macro)]

extern crate rocket;
extern crate ride_service;
extern crate diesel;

use self::ride_service::*;
use diesel::result::Error;
use ride_service::models::*;
use ride_service::dijkstra::*;
use rocket::{*, http::Status};
use rocket_contrib::json::Json;
use rocket::response::status;

use std::env;

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
pub fn create(mut new_drive: Json<NewDrive>) -> Result<status::Created<Json<Drive>>, Status> {
    let conn = establish_connection();

    new_drive.departure_location = new_drive.departure_location.trim().to_lowercase();
    new_drive.destination = new_drive.destination.trim().to_lowercase();

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

#[get("/advanced-search")]
pub fn advanced_search() -> Result<Json<Vec<Drive>>, Status> {
    let conn = establish_connection();
    

    let from: String = String::from("novi sad");
    let to: String = String::from("zajecar");

    let mut starting_ids: Vec<i32> = Vec::new();
    let mut ending_ids: Vec<i32> = Vec::new();

    let mut graph = Graph::new(Vec::new());

    let mut nodes: Vec<Node> = Vec::new();
    let drives = get_all(&conn).unwrap();

    let mut drive_inserted = false;
    for drive in &drives {
        for node in &mut nodes {
            if node.departure_location == drive.departure_location && node.destination == drive.destination && drive.free_places > 0 {
                node.distances.push(drive.distance as u64);
                node.departure_date_times.push(drive.departure_date_time as u64);
                node.planned_arival_times.push(drive.planned_arrival_time as u64);
                drive_inserted = true;
            }
        }
        if !drive_inserted {
            let mut new_node: Node = Node { 
                departure_location: drive.departure_location.clone(), 
                destination: drive.destination.clone(), 
                drive_ids: Vec::new(), 
                distances: Vec::new(), 
                departure_date_times: Vec::new(), 
                planned_arival_times: Vec::new(),
            };
            new_node.drive_ids.push(drive.id);
            new_node.distances.push(drive.distance as u64);
            new_node.departure_date_times.push(drive.departure_date_time as u64);
            new_node.planned_arival_times.push(drive.planned_arrival_time as u64);
            nodes.push(new_node);
        }
        
        drive_inserted = false;
    }

    for node1 in &nodes {
        for node2 in &nodes {
            if node1.destination == node2.departure_location && node1.departure_location.ne(&node2.destination) {
                
                // Ovde nadji minimum po vremenu i po duzini puta
                let mut min: u64 = std::u64::MAX;
                let mut min_i_idx: usize = std::usize::MAX;
                let mut min_j_idx: usize = std::usize::MAX;

                for i in 0..node1.departure_date_times.len() {
                    for j in 0..node2.departure_date_times.len() {
                        if node1.planned_arival_times[i] > node2.departure_date_times[j] {
                            continue;
                        }

                        let new_min: u64 = node1.distances[i];
                        if min > new_min {
                            min = new_min;
                            min_i_idx = i;
                            min_j_idx = j; 
                        }
                    }
                }

                if min != std::u64::MAX {

                    if node1.departure_location == from {
                        starting_ids.push(node1.drive_ids[min_i_idx])
                    }

                    if node2.destination == to {
                        ending_ids.push(node2.drive_ids[min_i_idx])
                    }

                    graph.connections.push(
                        dijkstra::Connection {                          
                            peers: (node1.drive_ids[min_i_idx], node2.drive_ids[min_j_idx]),
                            weight: min,
                        }
                    )
                    
                }
            }
        }
    }

    println!("Connections\n");
    for connection in &graph.connections {
        println!("{} to {} - weight = {}", connection.peers.0, connection.peers.1, connection.weight);
    } 
    println!("-----------------\n\n");

    dijkstra_fun_graph(&graph, &starting_ids, &ending_ids, &drives);

    // for ending_id in &ending_ids {
    //     println!("{}", *ending_id);
    // }


    Ok(Json(drives))
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
    // let conn = establish_connection();
    // _ = insert_d1(&conn);
    // _ = insert_d2(&conn);
    // _ = insert_d3(&conn);
    // _ = insert_d4(&conn);
    // _ = insert_d5(&conn);
    // _ = insert_d6(&conn);
    // _ = insert_d7(&conn);
    env::set_var("RUST_BACKTRACE", "1");
    // dijkstra_fun();

    rocket::ignite().mount("/api", routes![find, create, update, finish, reserve, delete, search, all, finished_driver, unfinished_driver, one_unfinished_driver,
                                                 one_finished, one_finished_driver, advanced_search],).launch();


}
