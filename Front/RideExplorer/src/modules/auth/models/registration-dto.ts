import { CarDTO } from "src/modules/shared/models/car-dto"

export interface RegistrationDTO {
    username: string
    password: string
    firstname: string
    lastname: string
    email: string
}

export interface DriverRegistrationDTO extends RegistrationDTO {
    car: CarDTO
}