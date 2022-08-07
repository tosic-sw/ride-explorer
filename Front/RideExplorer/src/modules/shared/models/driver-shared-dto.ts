import { CarDTO } from "./car-dto";

export interface DriverWithCarDTO {
    username: string
    password: string
    firstname: string
    lastname: string
    email: string
    car: CarDTO
}