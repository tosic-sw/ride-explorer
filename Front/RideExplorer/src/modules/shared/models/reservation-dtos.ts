export interface ReservationDTO {
    id: number 
    createdAt: number
    driveId: number
    passengerUsername: string
    driverUsername: string
    verified: boolean
}

export interface CreateReservationDTO {
    driveId: number
    driverUsername: string
}