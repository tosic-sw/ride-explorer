export interface NewDriveDTO {
    driver_username: string,
    departure_location: string,
    destination: string,
    departure_date_time: number,
    departure_address: string,
    free_places: number,
    planned_arrival_time: number,
    note: string,
    distance: number,
}

export interface DriveDTO {
    id: number,
    driver_username: string,
    departure_location: string,
    destination: string,
    departure_date_time: number,
    departure_address: string,
    free_places: number,
    planned_arrival_time: number,
    note: string,
    finished: boolean,
    distance: number,
}

export interface UpdateDriveDTO {
    id: number,
    departure_address: string,
    free_places: number,
    note: string,
}

export interface ReserveDTO {
    id: number,
    places: number,
}