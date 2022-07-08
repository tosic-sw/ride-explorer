table! {
    drives (id) {
        id -> Int4,
        departure_location -> Varchar,
        destination -> Varchar,
        departure_date_time -> Int8,
        departure_address -> Varchar,
        free_places -> Int4,
        planned_arrival_time -> Int8,
        note -> Varchar,
        finished -> Bool,
        distance -> Int4,
    }
}
