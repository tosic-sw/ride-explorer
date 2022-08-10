import { CarDTO } from "src/modules/shared/models/car-dto"

export interface UserDTO {
    username: string
    firstname: string
    lastname: string
    email: string
}

export interface DriverWithCarDTO extends UserDTO {
    car: CarDTO
}